mac:
	go env -w GOARCH=amd64
	go env -w GOOS=darwin
	go env -w CGO_ENABLED=0
	go env -w GO111MODULE=on
	go env -w GOPROXY=https://goproxy.cn,direct
	go mod  tidy

linux:
	go env -w GOARCH=amd64
	go env -w GOOS=linux
	go env -w CGO_ENABLED=0
	go env -w GO111MODULE=on
	go env -w GOPROXY=https://goproxy.cn,direct
	go mod  tidy

build-mac: mac
	go build -o gin-run-mac main.go

build-linux: linux
	go build -o gin-run-linux main.go

update:
	go mody tidy

run-debug:
	go run cmd/api/main.go --env=debug

run-release:
	go run cmd/api/main.go --env=release

server-kafka:
	go run server/kafkaServer.go --env=release

es-index:
	go run elasticsearch/index.go --env=release
