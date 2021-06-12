FROM golang:1.16

WORKDIR /go/src/app
COPY . .

RUN go get -d  ./...
RUN go install  ./...

RUN go get github.com/codegangsta/gin


EXPOSE 5000
EXPOSE 5100
