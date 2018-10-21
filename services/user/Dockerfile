FROM alpine:3.8
ARG service

ENV BINARY=$service

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN adduser -D -g '' appuser

RUN mkdir -p /go/bin/

COPY release/${service} /go/bin/
COPY config.json /go/bin/

WORKDIR /go/bin/

USER appuser

ENTRYPOINT /go/bin/$BINARY
