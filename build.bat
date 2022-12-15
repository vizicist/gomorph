go build gomorph.go
rm -fr bin
mkdir bin
mv gomorph.exe bin\gomorph.exe
copy SenselLib\winx64\*.dll bin
