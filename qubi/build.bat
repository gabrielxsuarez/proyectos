@echo off
echo Compilando qubi...

REM Verificar que Go esté instalado
go version >nul 2>&1
if %errorlevel% neq 0 (
    echo Error: Go no está instalado o no está en el PATH
    pause
    exit /b 1
)

REM Descargar dependencias
echo Descargando dependencias...
go mod download
if %errorlevel% neq 0 (
    echo Error al descargar dependencias
    pause
    exit /b 1
)

REM Instalar rsrc si no está disponible
echo Verificando herramienta rsrc...
rsrc -h >nul 2>&1
if %errorlevel% neq 0 (
    echo Instalando rsrc para agregar icono...
    go install github.com/akavel/rsrc@latest
    if %errorlevel% neq 0 (
        echo Error al instalar rsrc
        pause
        exit /b 1
    )
)

REM Generar archivo de recursos con icono
echo Generando recursos con icono...
rsrc -ico config\icono.ico -o rsrc.syso
if %errorlevel% neq 0 (
    echo Error al generar archivo de recursos
    pause
    exit /b 1
)

REM Verificar que GCC esté instalado (requerido para robotgo)
echo Verificando GCC para robotgo...
gcc --version >nul 2>&1
if %errorlevel% neq 0 (
    echo Advertencia: GCC no encontrado. Se requiere MinGW-w64 para el Mouse Jiggle.
    echo Continuando sin soporte para Mouse Jiggle...
)

REM Compilar el proyecto con CGO habilitado
echo Compilando proyecto...
set CGO_ENABLED=1
go build -ldflags "-H=windowsgui" -o qubi.exe .
if %errorlevel% neq 0 (
    echo Error durante la compilación
    pause
    exit /b 1
)

REM Limpiar archivo temporal
if exist rsrc.syso del rsrc.syso

echo Compilacion exitosa: Se ha generado qubi.exe
pause