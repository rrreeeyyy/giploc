VERSION := ${shell cat ./VERSION}
COMMIT = $(shell git describe --always)

GOOS = "linux"
GOARCH = "amd64"

build:
	GOOS=${GOOS} GOARCH=${GOARCH} go build -ldflags "-X github.com/rrreeeyyy/giploc/cmd.version=${VERSION}" -o giploc
