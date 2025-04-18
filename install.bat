@echo off


echo ==> Creating dist directory...
if not exist dist mkdir dist

echo ==> Changing to cmd directory for compilation...
cd cmd

echo ==> Compiling binary...
set GOOS=windows
go build -o ..\dist\commitia.exe ./

cd ..

if not exist dist\commitia.exe (
  echo Error: Binary 'dist\commitia.exe' was not generated.
  exit /b 1
)


echo ==> Installing binary...
if not exist "%USERPROFILE%\bin" mkdir "%USERPROFILE%\bin"
copy /Y "dist\commitia.exe" "%USERPROFILE%\bin\commitia.exe"


setx PATH "%PATH%;%USERPROFILE%\bin"

echo ==> Binary 'commitia.exe' successfully installed to %USERPROFILE%\bin\commitia.exe
echo Use 'commitia' to run the program.
echo Please restart your command prompt for the changes to take effect.
