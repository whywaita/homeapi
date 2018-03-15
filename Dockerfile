FROM golang:alpine

ENV GOPATH=/opt

RUN mkdir -p ${GOPATH}
ADD ./ /opt/src/github.com/whywaita/yayoi/

WORKDIR /opt/src/github.com/whywaita/yayoi/

RUN apk --update add git \
    && rm -rf /var/lib/apt/lists/* \
    && rm /var/cache/apk/* \
    && go get -u github.com/golang/dep/cmd/dep \
    && ${GOPATH}/bin/dep ensure -vendor-only \
    && go build -o /yayoi main.go

EXPOSE 8080

CMD ["/yayoi"]
