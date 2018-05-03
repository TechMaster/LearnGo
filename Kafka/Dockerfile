FROM golang:alpine

ARG LIBRDKAFKA_VERSION=0.11.4-r1

RUN apk update && \
	apk add --upgrade apk-tools && \
    apk add libressl --update-cache --repository http://nl.alpinelinux.org/alpine/edge/main && \
    apk add librdkafka=${LIBRDKAFKA_VERSION} --update-cache --repository http://nl.alpinelinux.org/alpine/edge/community && \
    apk add librdkafka-dev=${LIBRDKAFKA_VERSION} --update-cache --repository http://nl.alpinelinux.org/alpine/edge/community && \
    apk add git openssh openssl yajl-dev zlib-dev cyrus-sasl-dev openssl-dev build-base coreutils && \
    go get github.com/confluentinc/confluent-kafka-go/kafka