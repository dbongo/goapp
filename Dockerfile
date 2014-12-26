FROM dbongo/go

RUN mkdir -p /go/src/github.com/dbongo/goapp

ADD . /go/src/github.com/dbongo/goapp

WORKDIR /go/src/github.com/dbongo/goapp

RUN go get -d ./...
RUN go install ./...

EXPOSE 8000

ENTRYPOINT ["goapp"]
