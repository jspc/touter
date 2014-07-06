GO=go
GARGS=-a -v

all: install

install: touter

touter:
	$(GO) build $(GARGS) -o bin/touter src/touter.go

clean:
	rm -rfv bin/*
