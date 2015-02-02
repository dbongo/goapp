FROM google/golang

ADD . /gopath/src/github.com/dbongo/hackapp/
WORKDIR /gopath/src/github.com/dbongo/hackapp

RUN apt-get update
RUN make deps build install

EXPOSE 3000
ENTRYPOINT ["/usr/local/bin/hackappd"]
