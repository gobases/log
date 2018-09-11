@echo off
cd %~dp0
if %0 == "install" goto install



:install
echo -------------------------------------- begin to make install...
::go get github.com/chainlibs/gobtclib
mkdir %GOPATH%\src\github.com\gobasis
rmdir /s/q %GOPATH%\src\github.com\gobasis\log
mklink /D %GOPATH%\src\github.com\gobasis\log %cd%
echo -------------------------------------- finished successfully!
pause