os = linux
TAG ?= latest
SERVICE ?= salon

guard-%:
	@ if [ "${${*}}" = "" ]; then \
	echo "Environment variable $* not set"; \
	exit 1; \
	fi

all:	guard-SERVICE release build

.PHONY: release
release:	guard-SERVICE
	mkdir -p release
	CGO_ENABLED=0 GOOS=$(os) GOARCH=amd64 go build -o release/$(SERVICE)

push:	guard-SERVICE
	docker push amstee/easy-cut-$(SERVICE):$(TAG)

clean:
	rm -rf ./release

build:	guard-SERVICE
	docker build	--build-arg service=$(SERVICE) \
			-t amstee/easy-cut-$(SERVICE)":"$(TAG) .

