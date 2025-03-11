BINARY_NAME=secure-files-go-gui
BINARY_MACOS=${BINARY_NAME}-darwin_arm64
BINARY_LINUX=${BINARY_NAME}-linux_amd64
BINARY_WIN=${BINARY_NAME}-win_amd64.exe
OUT_PATH=out
BIN_PATH=${OUT_PATH}/bin

build:
	mkdir -p ${BIN_PATH}
	GOARCH=arm64 GOOS=darwin go build -o ${BIN_PATH}/${BINARY_MACOS} .
	GOARCH=amd64 GOOS=linux go build -o ${BIN_PATH}/${BINARY_LINUX} .
	GOARCH=amd64 GOOS=windows go build -o ${BIN_PATH}/${BINARY_WIN} .

clean:
	go clean
	rm -rf ${OUT_PATH}/*
	rm -rf test_files/*.enc