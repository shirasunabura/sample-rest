FROM google/cloud-sdk:alpine

LABEL maintainer="shirasu"

WORKDIR /go

ENV GOPATH=/go \
    PATH=/go/bin:/usr/local/go/bin:$PATH

ARG GOLANG_VERSION=1.14.15

RUN apk update && \
    apk add make gcc git musl-dev && \
    apk --no-cache add tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    \
    gcloud components install \
    app-engine-go && \
    \
    curl -o go.tgz -sSL "https://golang.org/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz" && \
    tar -C /usr/local -xzf go.tgz && \
    rm go.tgz && \
    \
    go get github.com/GoogleCloudPlatform/cloudsql-proxy/cmd/cloud_sql_proxy && \
    \
    mkdir /cloudsql && \
    \
    rm -rf src/*

ENTRYPOINT bash