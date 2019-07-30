ROOT=github.com/gopherdojo/dojo6/kadai2/pei
BIN=bin/imgconv
MAIN=cmd/cli.go
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
