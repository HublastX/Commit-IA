@echo off


echo ==> Creating dist directory...
if not exist dist mkdir dist

echo ==> Changing to cmd directory for compilation...
cd cmd

echo ==> Compiling binary...
set GOOS=windows
go build -o ..\dist\commitai.exe ./

cd ..

if not exist dist\commitai.exe (
  echo Error: Binary 'dist\commitai.exe' was not generated.
  exit /b 1
)


echo ==> Installing binary...
if not exist "%USERPROFILE%\bin" mkdir "%USERPROFILE%\bin"
copy /Y "dist\commitai.exe" "%USERPROFILE%\bin\commitai.exe"


setx PATH "%PATH%;%USERPROFILE%\bin"

echo ==> Binary 'commitai.exe' successfully installed to %USERPROFILE%\bin\commitai.exe
echo Use 'commitai' to run the program.
echo Please restart your command prompt for the changes to take effect.
