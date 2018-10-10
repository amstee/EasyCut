FROM golang:1.11-alpine3.8

RUN mkdir /service
WORKDIR /service

RUN go build -o build/rating .

CMD ["/service/build/rating"]