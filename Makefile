PROJECT=go-sitecheck
BIN=$(CURDIR)/bin
EXEC=$(PROJECT)

all: build

build: 	
	SET GOOS=linux& SET GOARCH=amd64& go build -o $(BIN)/$(EXEC) 
	SET GOOS=windows& SET GOARCH=amd64& go build -o $(BIN)/$(EXEC).exe

dep:
	go mod tidy
	