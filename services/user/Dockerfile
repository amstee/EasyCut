FROM alpine:3.8
ARG service
ARG api_id
ARG api_secret

ENV BINARY=$service
ENV API_CLIENT_ID=$api_id
ENV API_CLIENT_SECRET=$api_secret

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN adduser -D -g '' appuser

RUN mkdir -p /go/bin/

COPY release/${service} /go/bin/
COPY config.json /go/bin/

WORKDIR /go/bin/

USER appuser

ENTRYPOINT /go/bin/$BINARY
