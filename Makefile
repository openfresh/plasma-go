BASE_PACKAGE=github.com/openfresh/plasma-go
VERSION=$(shell git rev-parse --verify HEAD)

install-go:
		sh util/go.sh 1.8

install-glide:
		sh util/glide.sh v0.12.3

deps: install-glide
		glide install

deps-update: install-glide
		rm -rf ./vendor
		glide update

test:
		go test -v
