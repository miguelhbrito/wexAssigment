ENV = $(shell go env GOPATH)
GO_VERSION = $(shell go version)
GO111MODULE=on

config-up:
	echo "starting up configs"
	docker-compose up -d
config-down:
	docker-compose down