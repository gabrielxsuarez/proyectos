@echo off
echo ========================================
echo  Database Schema Extractor - Build
echo ========================================
echo.

echo Cleaning previous builds...
if exist schema-extractor.exe del schema-extractor.exe
if exist database-schema-extractor.exe del database-schema-extractor.exe

echo.
echo Downloading dependencies...
go mod tidy

echo.
echo Building application...
go build -o schema-extractor.exe cmd/main.go

if %ERRORLEVEL% EQU 0 (
    echo.
    echo ========================================
    echo  BUILD SUCCESSFUL!
    echo ========================================
    echo.
    echo Executable created: schema-extractor.exe
    echo.
    echo To use the application:
    echo 1. Create a config.yaml file with your database connections
    echo 2. Run: schema-extractor.exe
    echo.
    echo Example config.yaml:
    echo ifarmacia.fdb: "sysdba:masterkey@localhost:3050/C:\\path\\to\\ifarmacia.fdb"
    echo clientes.fdb: "sysdba:masterkey@localhost:3050/C:\\path\\to\\clientes.fdb"
    echo.
) else (
    echo.
    echo ========================================
    echo  BUILD FAILED!
    echo ========================================
    echo.
    echo Please check the error messages above.
    echo Make sure you have Go installed and properly configured.
    echo.
)

pause