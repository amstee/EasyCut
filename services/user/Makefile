os = linux
TAG ?= latest
SERVICE ?= user

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

clean:
	rm -rf ./release

build:	guard-SERVICE
	docker build	--build-arg service=$(SERVICE) \
			--build-arg api_id=$(API_CLIENT_ID) \
			--build-arg api_secret=$(API_CLIENT_SECRET) \
			-t amstee/easy-cut-$(SERVICE)":"$(TAG) .

