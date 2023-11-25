FROM golang:1.21.4-alpine3.18  as base
EXPOSE 2345 8090
RUN apk update && \
    apk upgrade && \
    apk add build-base git && \
    apk add nmap-ncat && \
    apk add curl && \
    apk add bind-tools && \
    apk add net-tools
RUN    go install github.com/githubnemo/CompileDaemon@latest  
RUN   CGO_ENABLED=0 go install github.com/go-delve/delve/cmd/api/dlv@latest