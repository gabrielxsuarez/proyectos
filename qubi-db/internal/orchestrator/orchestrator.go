package orchestrator

import (
	"database-schema-extractor/internal/config"
	"database-schema-extractor/internal/extractor"
	"database-schema-extractor/internal/extractor/firebird"
	"database-schema-extractor/internal/output"
	"fmt"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

// ExtractorOrchestrator coordinates parallel database schema extraction
type ExtractorOrchestrator struct {
	config      *config.Config
	maxWorkers  int
	timeout     time.Duration
	factories   []extractor.ExtractorFactory
	logger      *logrus.Entry
}

// NewExtractorOrchestrator creates a new orchestrator with default settings
func NewExtractorOrchestrator(cfg *config.Config, maxWorkers int) *ExtractorOrchestrator {
	return &ExtractorOrchestrator{
		config:     cfg,
		maxWorkers: maxWorkers,
		timeout:    30 * time.Second, // Default 30 second timeout
		factories:  []extractor.ExtractorFactory{firebird.NewFirebirdExtractorFactory()},
		logger:     logrus.WithField("component", "orchestrator"),
	}
}

// SetTimeout sets the timeout for database operations
func (eo *ExtractorOrchestrator) SetTimeout(timeout time.Duration) {
	eo.timeout = timeout
}

// AddExtractorFactory adds a new extractor factory to support additional database types
func (eo *ExtractorOrchestrator) AddExtractorFactory(factory extractor.ExtractorFactory) {
	eo.factories = append(eo.factories, factory)
}

// ExtractAll extracts schemas from all configured databases in parallel
func (eo *ExtractorOrchestrator) ExtractAll() []output.ExtractionResult {
	connections := eo.config.GetConnections()
	
	eo.logger.WithField("databases", len(connections)).Info("Starting parallel schema extraction")

	// Create channels for work distribution
	jobs := make(chan config.DatabaseConnection, len(connections))
	results := make(chan output.ExtractionResult, len(connections))

	// Start worker goroutines
	var wg sync.WaitGroup
	workerCount := eo.maxWorkers
	if workerCount > len(connections) {
		workerCount = len(connections)
	}

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go eo.worker(i, jobs, results, &wg)
	}

	// Send jobs to workers
	go func() {
		defer close(jobs)
		for _, conn := range connections {
			jobs <- conn
		}
	}()

	// Wait for all workers to complete
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	var extractionResults []output.ExtractionResult
	for result := range results {
		extractionResults = append(extractionResults, result)
	}

	eo.logSummary(extractionResults)
	return extractionResults
}

// worker processes database extraction jobs
func (eo *ExtractorOrchestrator) worker(id int, jobs <-chan config.DatabaseConnection, 
	results chan<- output.ExtractionResult, wg *sync.WaitGroup) {
	
	defer wg.Done()
	
	workerLogger := eo.logger.WithField("worker", id)
	workerLogger.Debug("Worker started")

	for conn := range jobs {
		workerLogger.WithField("database", conn.Name).Info("Processing database")
		
		result := eo.extractSingleDatabase(conn)
		results <- result
		
		if result.IsError() {
			workerLogger.WithError(result.Error).WithField("database", conn.Name).Error("Database extraction failed")
		} else {
			workerLogger.WithField("database", conn.Name).Info("Database extraction completed successfully")
		}
	}
	
	workerLogger.Debug("Worker finished")
}

// extractSingleDatabase extracts schema from a single database with timeout
func (eo *ExtractorOrchestrator) extractSingleDatabase(conn config.DatabaseConnection) output.ExtractionResult {
	// Create a channel to receive the result
	resultChan := make(chan output.ExtractionResult, 1)
	
	// Run extraction in a goroutine with timeout
	go func() {
		result := eo.performExtraction(conn)
		resultChan <- result
	}()

	// Wait for result or timeout
	select {
	case result := <-resultChan:
		return result
	case <-time.After(eo.timeout):
		return output.NewExtractionError(conn.Name, 
			fmt.Errorf("extraction timeout after %v", eo.timeout))
	}
}

// performExtraction performs the actual database extraction
func (eo *ExtractorOrchestrator) performExtraction(conn config.DatabaseConnection) output.ExtractionResult {
	// Find appropriate extractor factory
	var selectedExtractor extractor.DatabaseExtractor
	var err error

	for _, factory := range eo.factories {
		if factory.SupportsConnectionString(conn.ConnectionString) {
			selectedExtractor, err = factory.CreateExtractor(conn.ConnectionString)
			if err != nil {
				return output.NewExtractionError(conn.Name, 
					fmt.Errorf("failed to create extractor: %w", err))
			}
			break
		}
	}

	if selectedExtractor == nil {
		return output.NewExtractionError(conn.Name, 
			fmt.Errorf("no suitable extractor found for connection string"))
	}

	// Connect to database
	if err := selectedExtractor.Connect(conn.ConnectionString); err != nil {
		return output.NewExtractionError(conn.Name, 
			fmt.Errorf("failed to connect to database: %w", err))
	}
	defer selectedExtractor.Close()

	// Extract schema
	schema, err := selectedExtractor.ExtractSchema()
	if err != nil {
		return output.NewExtractionError(conn.Name, 
			fmt.Errorf("failed to extract schema: %w", err))
	}

	return output.NewExtractionResult(conn.Name, schema)
}

// logSummary logs a summary of extraction results
func (eo *ExtractorOrchestrator) logSummary(results []output.ExtractionResult) {
	successful := 0
	failed := 0

	for _, result := range results {
		if result.IsSuccess() {
			successful++
		} else {
			failed++
		}
	}

	eo.logger.WithFields(logrus.Fields{
		"total":      len(results),
		"successful": successful,
		"failed":     failed,
	}).Info("Schema extraction completed")
}