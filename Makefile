GO=go
GARGS=-a -v
PREFIX=/usr/local

all: build

build: touter

touter:
	$(GO) build $(GARGS) -o bin/touter src/touter.go

clean:
	rm -rfv bin/*

install:
	install bin/touter $(PREFIX)/sbin/touter
	install sample_profiles.ini /etc/touter.ini
