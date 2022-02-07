GOARCH := amd64

all: mac windows linux

mac:
	env CGO_ENABLED=0 GOARCH=${GOARCH} GOOS=darwin go build -o ./build/mac/ghanamps -ldflags "-X 'main.version=$$(git tag -l | tail -n1)' -X 'main.arch=darwin/${GOARCH}'" ./cmd/ghanamps

windows:
	env CGO_ENABLED=0 GOARCH=${GOARCH} GOOS=windows go build -o ./build/windows/ghanamps.exe -ldflags "-X 'main.version=$$(git tag -l | tail -n1)' -X 'main.arch=windows/${GOARCH}'" ./cmd/ghanamps

linux:
	env CGO_ENABLED=0 GOARCH=${GOARCH} GOOS=linux go build -o ./build/linux/ghanamps -ldflags "-X 'main.version=$$(git tag -l | tail -n1)' -X 'main.arch=linux/${GOARCH}'" ./cmd/ghanamps


dist: all
	zip -mj ghanamps-mac.zip ./build/mac/ghanamps
	zip -mj ghanamps-linux.zip ./build/linux/ghanamps
	zip -mj ghanamps-windows.zip ./build/windows/ghanamps.exe
