default: build_win build_lin build_osx

build_win:
	GOOS=windows GOARCH=amd64 go build -o http2push-win32.exe -ldflags="-s -w" .

build_lin:
	GOOS=linux GOARCH=amd64 go build -o http2push-linux-amd64 -ldflags="-s -w" .
	GOOS=linux GOARCH=arm go build -o http2push-linux-arm32 -ldflags="-s -w" .
	GOOS=linux GOARCH=arm64 go build -o http2push-linux-arm64 -ldflags="-s -w" .

build_osx:
	GOOS=darwin GOARCH=amd64 go build -o http2push-osx-amd64 -ldflags="-s -w" .