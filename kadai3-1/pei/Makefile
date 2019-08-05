ROOT=github.com/gopherdojo/dojo6/kadai3-1/pei
BIN=bin/typing
MAIN=main.go
TEST=...

.PHONY: build
build: ${MAIN}
	go build -o ${BIN} ${GOPATH}/src/${ROOT}/$?

.PHONY: test
test:
	go test -v -cover ${ROOT}/${TEST}

.PHONY: clean
clean:
	rm ${BIN}
	go clean
