BINARY=fj-server
GO_FILES=$(wildcard *.go)


build: $(BINARY)

$(BINARY): $(GO_FILES)
	go build -o $(BINARY) $(GO_FILES)