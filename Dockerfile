FROM dbongo/go

RUN mkdir -p /go/src/github.com/dbongo/hackapp
ADD . /go/src/github.com/dbongo/hackapp
WORKDIR /go/src/github.com/dbongo/hackapp

RUN go get -d ./...
RUN go install ./...

EXPOSE 8000
ENTRYPOINT ["hackapp"]
