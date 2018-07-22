FROM golang:1.9.4

ENV GOPATH $GOPATH:/go
ENV DB_SOURCE="root:root@tcp(docker.for.mac.localhost:3306)/auth_server?charset=utf8&parseTime=True"

RUN go get github.com/golang/dep/cmd/dep
RUN go get github.com/revel/cmd/revel
RUN go get github.com/tomoyane/revel-performance

RUN mkdir -p /go/src/github.com/tomoyane/revel-performance
COPY . /go/src/github.com/tomoyane/revel-performance

WORKDIR /go/src/github.com/tomoyane/revel-performance

RUN dep ensure

CMD revel run