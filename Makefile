PROJECT=npm-binary-export.kit
OUTPUT=bin/$(PROJECT)
VERSION=0.0.1

.PHONY: build
build:
	go mod vendor
	CGO_ENABLED=0 go build -ldflags="-s" -mod=vendor -trimpath -o $(OUTPUT)-$(VERSION)-darwin

clean:
	go clean
	rm -f $(OUTPUT)
	rm -rf vendor/

run: build
	$(OUTPUT)

build-linux:
	go mod vendor
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s" -mod=vendor -trimpath -o $(OUTPUT)-$(VERSION)-linux-amd64

# 精简版程序
build-linux-upx:
	go mod vendor
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s" -mod=vendor -trimpath -o $(OUTPUT)-$(VERSION)-linux-amd64
	upx -6 $(OUTPUT)-$(VERSION)-linux-amd64


build-mac:
	go mod vendor
	CGO_ENABLED=0 go build -ldflags="-s" -mod=vendor -trimpath -o $(OUTPUT)-$(VERSION)-darwin

build-windows:
	go mod vendor
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-s" -mod=vendor -trimpath -o $(OUTPUT)-$(VERSION).exe
	upx -6 $(OUTPUT).exe