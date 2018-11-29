FROM golang:1.10-alpine3.8 as builder
LABEL maintainer="v.zorin@anchorfree.com"

RUN apk add --no-cache git 
COPY . /go/src/github.com/anchorfree/diffence
RUN cd /go && go build github.com/anchorfree/diffence/cmd/diffence

FROM alpine:3.8
LABEL maintainer="v.zorin@anchorfree.com"

RUN apk add --no-cache git bash && wget -q -O /etc/apk/keys/sgerrand.rsa.pub https://raw.githubusercontent.com/sgerrand/alpine-pkg-git-crypt/master/sgerrand.rsa.pub \
    && wget https://github.com/sgerrand/alpine-pkg-git-crypt/releases/download/0.6.0-r0/git-crypt-0.6.0-r0.apk \
    && apk add git-crypt-0.6.0-r0.apk \
    && rm git-crypt-0.6.0-r0.apk

COPY --from=builder /go/diffence /usr/local/bin/diffence

ENTRYPOINT ["/bin/bash"]
