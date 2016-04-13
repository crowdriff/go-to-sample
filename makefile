version=1.0.1

.PHONY: all

all:
	@echo "make <cmd>"
	@echo ""
	@echo "commands:"
	@echo "  build         - build the dist binary"
	@echo "  clean         - clean the dist build"
	@echo "  run           - run local dev version"
	@echo "  test          - standard go test"
	@echo ""
	@echo "  tools         - go gets a bunch of tools for dev"
	@echo "  deps          - pull and setup dependencies"
	@echo "  update_deps   - update deps lock file"

build: clean
	@go vet ./...
	@golint ./...
	@go build -o ./bin/go-to-sample.bin -ldflags "-X main.version=$(version)" main.go

clean:
	@rm -rf ./bin

deps:
	@glock sync -n github.com/crowdriff/go-to-sample < Glockfile

run:
	@fresh

test:
	@ginkgo -r -v -cover

tools:
	go get github.com/robfig/glock
	go get github.com/pilu/fresh
	go get github.com/golang/lint/golint
	go get github.com/onsi/ginkgo/ginkgo
	go get github.com/onsi/gomega

update_deps:
	@glock save -n github.com/crowdriff/go-to-sample > Glockfile
