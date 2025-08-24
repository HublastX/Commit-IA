@echo off
echo ==> Creating dist directory...
if not exist dist mkdir dist

echo ==> Changing to cmd directory...
cd cmd

echo ==> Compiling Windows binary...
set GOOS=windows
set GOARCH=amd64
go build -o ..\dist\commitai.exe ./

cd ..

if not exist dist\commitai.exe (
    echo Error: Binary 'dist\commitai.exe' was not generated.
    exit /b 1
)

echo ==> Windows build completed: dist\commitai.exe
