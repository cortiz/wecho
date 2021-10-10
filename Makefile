build:
	go build -o bin/wecho wecho.go

run:
	go run wecho.go


compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o bin/wecho-arm wecho.go
	GOOS=linux GOARCH=arm64 go build -o bin/wecho-linux-arm64 wecho.go
	GOOS=linux GOARCH=amd64 go build -o bin/wecho-linux-amd64 wecho.go
	GOOS=windows GOARCH=amd64 go build -o bin/wecho-windows-amd64 wecho.go 
	GOOS=darwin  GOARCH=amd64 go build -o bin/wecho-osx-intel wecho.go 
	GOOS=darwin GOARCH=arm64 go build -o bin/wecho-osx-m1 wecho.go 
all: build
