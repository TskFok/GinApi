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
	go build -o gin-run-mac -ldflags "-w -s"  -trimpath ./bin/api/main.go

build-linux: linux
	go build -o gin-run-linux -ldflags "-w -s"  -trimpath ./bin/api/main.go

update:
	go mody tidy

run-debug: mac
	go run bin/api/main.go --env=debug

run-release: mac
	go run bin/api/main.go --env=release

server-kafka: mac
	go run bin/server/kafkaServer.go --env=release

build-server-kafka-mac: mac
	go build -o gin-kafka-run-mac -ldflags "-w -s"  -trimpath ./bin/server/kafkaServer.go

build-server-kafka-linux: linux
	go build -o gin-kafka-run-linux -ldflags "-w -s"  -trimpath ./bin/server/kafkaServer.go

es-index: mac
	go run bin/cli/main.go es:index --env=debug

build-cli-mac: mac
	go build -o gin-cli-run-mac -ldflags "-w -s"  -trimpath ./bin/cli/main.go

build-cli-linux: linux
	go build -o gin-cli-run-linux -ldflags "-w -s"  -trimpath ./bin/cli/main.go