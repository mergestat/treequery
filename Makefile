.PHONY: all clean

all: clean tq

clean:
	-rm -f tq

tq:
	go build -o $@ .
