@echo off
echo ==> Creating dist directory...
if not exist dist mkdir dist

echo ==> Changing to cmd directory...
cd cmd

echo ==> Compiling Windows binary...
set GOOS=windows
set GOARCH=amd64
go build -o ..\dist\commitia.exe ./

cd ..

if not exist dist\commitia.exe (
    echo Error: Binary 'dist\commitia.exe' was not generated.
    exit /b 1
)

echo ==> Windows build completed: dist\commitia.exe
