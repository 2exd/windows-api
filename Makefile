# Makefile for cross-compiling client and server executables
# Compiler and linker options
GO = go
GOBUILD = $(GO) build
GOCLEAN = $(GO) clean

# Directories
BIN_DIR = ./bin
CLIENT_DIR = client
SERVER_DIR = server

# Executables
CLIENT_EXEC = $(BIN_DIR)/client/
SERVER_EXEC = $(BIN_DIR)/server/

# OS and ARCH combinations for cross-compilation
OS_ARCH = "windows/386 windows/amd64 darwin/amd64 linux/386 linux/amd64"

# 编译到 Linux
.PHONY: build-linux
build-linux:
     GOOS=linux GOARCH=amd64 ${GOBUILD} -o ${CLIENT_EXEC}/client-linux $(CLIENT_DIR)/client.go

# 编译到 macOS
.PHONY: build-darwin
build-darwin:
    CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 ${GOBUILD} -o ${CLIENT_EXEC}/client-darwin $(CLIENT_DIR)/client.go

# 编译到 windows
.PHONY: build-windows
build-windows:
	CGO_ENABLED=1
	GOOS="windows"
	GOARCH="amd64"
	go build -o .\bin\client\ .\client\client.go
#    CGO_ENABLED=0 GOOS=windows GOARCH=amd64 ${GOBUILD} -o ${CLIENT_EXEC}/client-windows.exe $(CLIENT_DIR)/client.go

# 编译到 全部平台
.PHONY: build-all
build-all:
	make clean
	mkdir -p ${CLIENT_EXEC} ${SERVER_EXEC}
	make build-linux
	make build-darwin
	make build-windows

clean:
	@echo "Cleaning..."
	$(GOCLEAN)
	del /Q $(BIN_DIR)/client/* $(BIN_DIR)/server/*
	#rm -f $(BIN_DIR)/client/* $(BIN_DIR)/server/*


.PHONY: all clean
