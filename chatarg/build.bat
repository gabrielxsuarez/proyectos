@echo off
echo 🚀 Compilando ChatArg...
echo =============================

REM Compilar con flags para reducir detección de antivirus
go build -ldflags="-s -w" -o chatarg.exe main.go struct_envios.go struct_recibos.go

if %ERRORLEVEL% EQU 0 (
    echo ✅ Compilación exitosa: chatarg.exe
    echo 📦 Tamaño del archivo:
    dir /B chatarg.exe | findstr chatarg.exe >nul && for %%F in (chatarg.exe) do echo    %%~zF bytes
) else (
    echo ❌ Error en la compilación
    pause
    exit /b 1
)

echo.
echo 💡 Flags utilizados:
echo    -ldflags="-s -w" = Elimina símbolos y tabla de debug
echo    -o chatarg.exe = Nombre del ejecutable de salida
echo.
echo 🛡️ Si el antivirus lo detecta, agregar excepción en Windows Defender
pause