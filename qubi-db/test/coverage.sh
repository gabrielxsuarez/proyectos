#!/bin/bash

# Script to run tests with coverage analysis

echo "Running Database Schema Extractor Tests with Coverage"
echo "======================================================"

# Create coverage directory
mkdir -p coverage

# Run tests with coverage for each package
echo "Running tests with coverage..."

# Test config package
go test -coverprofile=coverage/config.out ./test/config
go tool cover -html=coverage/config.out -o coverage/config.html

# Test schema package  
go test -coverprofile=coverage/schema.out ./test/schema
go tool cover -html=coverage/schema.out -o coverage/schema.html

# Test output package
go test -coverprofile=coverage/output.out ./test/output  
go tool cover -html=coverage/output.out -o coverage/output.html

# Test firebird extractor package
go test -coverprofile=coverage/firebird.out ./test/extractor/firebird
go tool cover -html=coverage/firebird.out -o coverage/firebird.html

# Test orchestrator package
go test -coverprofile=coverage/orchestrator.out ./test/orchestrator
go tool cover -html=coverage/orchestrator.out -o coverage/orchestrator.html

# Generate combined coverage report
echo "mode: set" > coverage/combined.out
tail -n +2 coverage/config.out >> coverage/combined.out
tail -n +2 coverage/schema.out >> coverage/combined.out  
tail -n +2 coverage/output.out >> coverage/combined.out
tail -n +2 coverage/firebird.out >> coverage/combined.out
tail -n +2 coverage/orchestrator.out >> coverage/combined.out

# Generate combined HTML report
go tool cover -html=coverage/combined.out -o coverage/combined.html

# Show coverage summary
echo ""
echo "Coverage Summary:"
echo "=================="
go tool cover -func=coverage/combined.out

echo ""
echo "Coverage reports generated in coverage/ directory:"
echo "- coverage/combined.html (overall coverage)"
echo "- coverage/config.html"
echo "- coverage/schema.html" 
echo "- coverage/output.html"
echo "- coverage/firebird.html"
echo "- coverage/orchestrator.html"