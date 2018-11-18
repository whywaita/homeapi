FROM golang:alpine

ENV GOPATH=/opt

RUN mkdir -p ${GOPATH}

WORKDIR /opt/src/github.com/whywaita/yayoi/
ADD Gopkg.lock /opt/src/github.com/whywaita/yayoi/Gopkg.lock
ADD Gopkg.toml /opt/src/github.com/whywaita/yayoi/Gopkg.toml

RUN apk --update add git \
    && rm -rf /var/lib/apt/lists/* \
    && rm /var/cache/apk/* \
    && go get -u github.com/golang/dep/cmd/dep \
    && ${GOPATH}/bin/dep ensure -vendor-only 

ADD ./ /opt/src/github.com/whywaita/yayoi/
RUN go build -o /yayoi main.go

EXPOSE 8080

CMD ["/yayoi"]
