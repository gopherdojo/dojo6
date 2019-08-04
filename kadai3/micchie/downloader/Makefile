NAME := typing-game
SRCS := $(shell find ./ -type f -name '*.go')
VERSION := v0.0.1
REVISION := $(shell git rev-parse --short HEAD)
PKGPATH := github.com/gopherdojo/dojo6/kadai3/micchie/typing-game
LDFLAGS := -ldflags \
	'-s -w \
	-X "$(PKGPATH)/cmd.Version=$(VERSION)" \
	-X "$(PKGPATH)/cmd.Revision=$(REVISION)" \
	-extldflags "-static"'

DIST_DIRS := find * -type d -exec

$(NAME): $(SRCS)
	go build -a -tags netgo -installsuffix netgo $(LDFLAGS) -o $(NAME)

.PHONY: install
install:
	go install $(LDFLAGS)

.PHONY: clean
clean:
	rm -rf $(NAME)

.PHONY: cross-build
cross-build: deps
	for os in darwin linux windows; do \
		for arch in amd64 386; do \
			GOOS=$$os GOARCH=$$arch CGO_ENABLED=0 go build -a -tags netgo -installsuffix netgo $(LDFLAGS) -o dist/$$os-$$arch/$(NAME); \
		done; \
	done

.PHONY: dist
dist:
	cd dist && \
  $(DIST_DIRS) cp ../LICENSE {} \; && \
  $(DIST_DIRS) cp ../README.md {} \; && \
  $(DIST_DIRS) tar -zcf $(NAME)-$(VERSION)-{}.tar.gz {} \; && \
  $(DIST_DIRS) zip -r $(NAME)-$(VERSION)-{}.zip {} \; && \
  cd ..
