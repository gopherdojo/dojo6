ROOT=github.com/gopherdojo/dojo6/kadai4/pei
BIN=fortune
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
