all: mac windows linux

mac:
	env CGO_ENABLED=0 GOARCH=amd64 GOOS=darwin go build -o ./build/mac/ghanamps ./cmd/ghanamps

windows:
	env CGO_ENABLED=0 GOARCH=amd64 GOOS=windows go build -o ./build/windows/ghanamps.exe ./cmd/ghanamps

linux:
	env CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o ./build/linux/ghanamps ./cmd/ghanamps

dist: all
	zip -mj ghanamps-mac.zip ./build/mac/ghanamps
	zip -mj ghanamps-linux.zip ./build/linux/ghanamps
	zip -mj ghanamps-windows.zip ./build/windows/ghanamps.exe
