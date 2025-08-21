package main

import (
	"database-schema-extractor/internal/config"
	"database-schema-extractor/internal/orchestrator"
	"database-schema-extractor/internal/output"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

const (
	ExitSuccess           = 0
	ExitConfigError       = 1
	ExitExtractionError   = 2
	ExitOutputError       = 3
	ExitPartialFailure    = 4
)

func main() {
	exitCode := run()
	os.Exit(exitCode)
}

func run() int {
	// Setup logging
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.InfoLevel)

	logger := logrus.WithField("component", "main")
	logger.Info("Starting Database Schema Extractor")

	// Load configuration
	configManager := config.NewConfigManager()
	cfg, err := configManager.LoadConfig()
	if err != nil {
		logger.WithError(err).Error("Failed to load configuration")
		
		// Provide helpful error messages based on error type
		switch {
		case err == config.ErrConfigFileNotFound:
			fmt.Fprintf(os.Stderr, "Error: config.yaml file not found in current directory\n")
			fmt.Fprintf(os.Stderr, "Please create a config.yaml file with your database connections\n")
			fmt.Fprintf(os.Stderr, "Example:\n")
			fmt.Fprintf(os.Stderr, "  ifarmacia.fdb: \"sysdba:masterkey@localhost:3050/path/to/ifarmacia.fdb\"\n")
			fmt.Fprintf(os.Stderr, "  clientes.fdb: \"sysdba:masterkey@localhost:3050/path/to/clientes.fdb\"\n")
		case err == config.ErrNoDatabasesConfigured:
			fmt.Fprintf(os.Stderr, "Error: No databases configured in config.yaml\n")
			fmt.Fprintf(os.Stderr, "Please add at least one database connection to your config.yaml file\n")
		default:
			fmt.Fprintf(os.Stderr, "Configuration error: %v\n", err)
		}
		
		return ExitConfigError
	}

	logger.WithField("databases", len(cfg.Databases)).Info("Configuration loaded successfully")

	// Create orchestrator and extract schemas
	orch := orchestrator.NewExtractorOrchestrator(cfg, 5) // 5 concurrent workers
	results := orch.ExtractAll()

	if len(results) == 0 {
		logger.Warn("No databases were processed")
		return ExitSuccess
	}

	// Generate output files and track results
	generator := output.NewYAMLGenerator()
	successCount := 0
	extractionErrors := 0
	outputErrors := 0

	for _, result := range results {
		if result.IsError() {
			extractionErrors++
			logger.WithError(result.Error).WithField("database", result.DatabaseName).Error("Database extraction failed")
		} else {
			successCount++
		}

		if err := generator.GenerateOutput(result); err != nil {
			outputErrors++
			logger.WithError(err).WithField("database", result.DatabaseName).Error("Failed to generate output file")
		}
	}

	// Log final summary
	logger.WithFields(logrus.Fields{
		"total":             len(results),
		"successful":        successCount,
		"extraction_errors": extractionErrors,
		"output_errors":     outputErrors,
	}).Info("Database Schema Extractor completed")

	// Determine exit code based on results
	if outputErrors > 0 {
		fmt.Fprintf(os.Stderr, "Warning: %d output file generation errors occurred\n", outputErrors)
		return ExitOutputError
	}

	if extractionErrors > 0 {
		if successCount > 0 {
			fmt.Fprintf(os.Stderr, "Warning: %d database extraction errors occurred, but %d succeeded\n", 
				extractionErrors, successCount)
			return ExitPartialFailure
		} else {
			fmt.Fprintf(os.Stderr, "Error: All database extractions failed\n")
			return ExitExtractionError
		}
	}

	fmt.Printf("Successfully extracted schemas from %d database(s)\n", successCount)
	return ExitSuccess
}