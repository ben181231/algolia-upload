FROM golang:1.13 AS building

COPY . /usr/src/app

WORKDIR /usr/src/app

RUN \
  CGO_ENABLED=0 go build -o /usr/bin/algolia-upload .

FROM alpine:3.9

COPY --from=building /usr/bin/algolia-upload /usr/bin/algolia-upload

CMD "algolia-upload"
