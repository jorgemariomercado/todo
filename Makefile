

PROJ := todo
ORG_PATH := git.local/jmercado
REPO_PATH := $(ORG_PATH)/$(PROJ)
GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)

export GOBIN=$(PWD)/bin

ifeq ($(strip $(VERSION)),)
	VERSION := $(GIT_BRANCH)
endif


LD_FLAGS="-w -X main.version=$(VERSION)"

.PHONY: default


.PHONY: default
default: clean build


.PHONY: clean
clean:
	rm -rf ./bin/*

.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v


# Run quality control checks
# go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
# go run golang.org/x/vuln/cmd/govulncheck@latest ./...

.PHONY: test
test:
	go mod verify
	go vet ./...
	go test -race -buildvcs -vet=off ./...


.PHONY: build
build: tidy
	mkdir -p bin/
	go install -v -ldflags $(LD_FLAGS) $(REPO_PATH)/cmd/$(PROJ)


.PHONY: api
api:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    api/*.proto

.PHONY: dev
dev:
	./bin/todo serve config.yml

.PHONY: run
run:
	./bin/todo serve config.yml


.PHONY: web-install
web-install:
	yarn install

