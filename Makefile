build: install
	go build -o make/zscnetworksupport-api
install:
	go install

clean:
	rm -rf make/zscnetworksupport-api

update:


help:

default: build
