FROM golang:1.10.0

RUN go get github.com/golang/dep/cmd/dep && \
    go get github.com/revel/cmd/revel && \
    go get github.com/tomoyane/revel-performance/app/tmp

ENV GOPATH="/go"
ENV PATH="$PATH:$GOPATH/bin"
ENV DB_SOURCE="root:root@tcp(docker.for.mac.localhost:3306)/auth_server?charset=utf8&parseTime=True"

WORKDIR src/github.com/tomoyane/revel-performance
RUN dep ensure

CMD revel run dev