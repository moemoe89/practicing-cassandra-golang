FROM golang:latest

RUN mkdir -p /go/src/github.com/moemoe89/practicing-cassandra-golang

WORKDIR /go/src/github.com/moemoe89/practicing-cassandra-golang

COPY . /go/src/github.com/moemoe89/practicing-cassandra-golang

RUN go mod download
RUN go install

ENTRYPOINT /go/bin/practicing-cassandra-golang
