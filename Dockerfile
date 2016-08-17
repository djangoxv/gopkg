FROM golang:1.6.3
ADD ./gopkg /go/bin/gopkg
CMD /go/bin/gopkg
EXPOSE 8080
