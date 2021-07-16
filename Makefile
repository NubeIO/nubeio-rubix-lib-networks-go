
HTTP_ADDRESS := 127.0.0.1:1920
DATABASE_URI := ../../ugin.db
DATABASE_ADAPTER := sqlite3

export  HTTP_ADDRESS DATABASE_URI DATABASE_ADAPTER

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=mybinary
BINARY_UNIX=$(BINARY_NAME)_unix
BINARY_BBB=$(BINARY_NAME)_bbb

curl-networks:
	curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET http://$(HTTP_ADDRESS)/api/networks

sqlite-networks:
	sqlite3 $(DATABASE_URI) -header -column -echo 'select * from networks;'


all: test build
build:
		$(GOBUILD) -o $(BINARY_UNIX) -v
test:
		$(GOTEST) -v ./...
clean:
		$(GOCLEAN)
		rm -f $(BINARY_NAME)
		rm -f $(BINARY_UNIX)
run:
		$(GOBUILD) -o $(BINARY_NAME) -v ./...
		./$(BINARY_NAME)

docs-swag:
	swag init

# Cross compilation
build-linux:
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
build-bbb:
		GOOS=linux GOARCH=arm CGO_ENABLED=1 CC=arm-linux-gnueabi-gcc GOARM=7 go build -ldflags="-extldflags=-static" -tags sqlite_omit_load_extension -o $(BINARY_BBB) -v

