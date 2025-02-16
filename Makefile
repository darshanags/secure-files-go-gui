BINARY_NAME=secure-files-go-gui

build:
	mkdir -p out/bin
	GOARCH=arm64 GOOS=darwin go build -o out/bin/${BINARY_NAME}-darwin_arm64 .
	GOARCH=amd64 GOOS=linux go build -o out/bin/${BINARY_NAME}-linux_amd64 .
	GOARCH=amd64 GOOS=windows go build -o out/bin/${BINARY_NAME}-win_amd64.exe .

# run: build
# 	./out/bin/${BINARY_NAME}-darwin_arm64

clean:
	go clean
	rm -rf out/bin/${BINARY_NAME}-darwin_arm64
	rm -rf out/bin/${BINARY_NAME}-linux_amd64
	rm -rf out/bin/${BINARY_NAME}-win_amd64.exe