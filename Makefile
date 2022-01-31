all: mac windows linux

mac:
	env GOARCH=amd64 GOOS=darwin go build -o ./build/mac/ghanamps

windows:
	env GOARCH=amd64 GOOS=windows go build -o ./build/windows/ghanamps.exe

linux:
	env GOARCH=amd64 GOOS=linux go build -o ./build/linux/ghanamps

dist: all
	cd ./build/mac && zip -m ghanamps-mac.zip ghanamps
	cd ./build/linux && zip -m ghanamps-linux.zip ghanamps
	cd ./build/windows && zip -m ghanamps-windows.zip ghanamps.exe
