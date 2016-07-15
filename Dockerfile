FROM golang:1.6.2
ADD ./gopkg /go/bin/gopkg
ENTRYPOINT /go/bin/gopkg
EXPOSE 8080
