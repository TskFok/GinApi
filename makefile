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
	go build -o gin-run-mac -ldflags "-w -s"  -trimpath ./cmd/api/main.go

build-linux: linux
	go build -o gin-run-linux -ldflags "-w -s"  -trimpath cmd/api/main.go

update:
	go mody tidy

run-debug:
	go run cmd/api/main.go --env=debug

run-release:
	go run cmd/api/main.go --env=release

server-kafka:
	go run server/kafkaServer.go --env=release

build-server-kafka-mac: mac
	go build -o gin-kafka-run-mac -ldflags "-w -s"  -trimpath ./server/kafkaServer.go

build-server-kafka-linux: linux
	go build -o gin-kafka-run-linux -ldflags "-w -s"  -trimpath ./server/kafkaServer.go

es-index:
	go run elasticsearch/index.go --env=release
