FROM alpine:latest

RUN apk update && apk add \
    curl \
    bind-tools \
    net-tools \
    nmap-ncat \
	bash \
    && rm -rf /var/cache/apk/*

CMD ["/bin/bash"]