FROM golang:1.7.4-wheezy

RUN apt-get update && apt-get -yq install libgeoip-dev

ADD data /data

ADD . /go/src/github.com/transitorykris/geoful
WORKDIR /go/src/github.com/transitorykris/geoful
RUN go get
RUN go build -o /app

EXPOSE 8080

CMD ["/app"]
