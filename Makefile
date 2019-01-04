all: bin/gedcom5

bin/gedcom5: bin $(shell find . -name '*.go')
	cd cmd/gedcom5 && go build -o ../../bin/gedcom5

test:
	go test -v -coverprofile=coverage.out ./... && go tool cover -o coverage.html -html=coverage.out

bin:
	mkdir -p bin

generated:
	rm -f *_gen_*.go
	go generate ./...

clean:
	rm -rf bin

.PHONY: test
.PHONY: all
.PHONY: clean
.PHONY: generated
