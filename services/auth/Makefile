os = linux
TAG ?= latest
SERVICE ?= auth

guard-%:
	@ if [ "${${*}}" = "" ]; then \
	echo "Environment variable $* not set"; \
	exit 1; \
	fi

all:	guard-SERVICE release build

.PHONY: release
release:	guard-SERVICE
	mkdir -p release
	GOOS=$(os) GOARCH=amd64 go build -o release/$(SERVICE)

clean:
	rm -rf ./release

build:	guard-SERVICE
	docker build --build-arg service=$(SERVICE) -t amstee/easy-cut-$(SERVICE)":"$(TAG) .

