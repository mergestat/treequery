.PHONY: all clean test vet

all: clean tq

clean:
	-rm -f tq

tq:
	go build -o $@ .

test:
	go test -v ./...

vet:
	go vet -v ./...
