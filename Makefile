all: bin/gedcom5

bin/gedcom5: bin $(shell find . -name '*.go')
	cd cmd/gedcom5 && go build -o ../../bin/gedcom5

test:
	go test -v ./...

bin:
	mkdir -p bin

clean:
	rm -rf bin

.PHONY: test
.PHONY: all
.PHONY: clean
