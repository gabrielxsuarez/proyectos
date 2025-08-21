package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Test runner to execute all unit tests
func main() {
	fmt.Println("Running Database Schema Extractor Test Suite")
	fmt.Println("=" + strings.Repeat("=", 50))

	// Get the project root directory
	projectRoot, err := getProjectRoot()
	if err != nil {
		fmt.Printf("Error finding project root: %v\n", err)
		os.Exit(1)
	}

	// Test packages to run
	testPackages := []string{
		"./test/config",
		"./test/schema",
		"./test/output",
		"./test/extractor/firebird",
		"./test/orchestrator",
	}

	totalTests := 0
	passedTests := 0
	failedPackages := []string{}

	for _, pkg := range testPackages {
		fmt.Printf("\nRunning tests for %s...\n", pkg)
		fmt.Println("-" + strings.Repeat("-", 30))

		cmd := exec.Command("go", "test", "-v", pkg)
		cmd.Dir = projectRoot
		output, err := cmd.CombinedOutput()

		fmt.Print(string(output))

		if err != nil {
			fmt.Printf("❌ Tests failed for %s\n", pkg)
			failedPackages = append(failedPackages, pkg)
		} else {
			fmt.Printf("✅ Tests passed for %s\n", pkg)
			passedTests++
		}
		totalTests++
	}

	// Print summary
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("TEST SUMMARY")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Printf("Total packages tested: %d\n", totalTests)
	fmt.Printf("Passed: %d\n", passedTests)
	fmt.Printf("Failed: %d\n", len(failedPackages))

	if len(failedPackages) > 0 {
		fmt.Println("\nFailed packages:")
		for _, pkg := range failedPackages {
			fmt.Printf("  - %s\n", pkg)
		}
		os.Exit(1)
	} else {
		fmt.Println("\n🎉 All tests passed!")
	}
}

func getProjectRoot() (string, error) {
	// Start from current directory and walk up to find go.mod
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}

	return "", fmt.Errorf("could not find project root (go.mod not found)")
}