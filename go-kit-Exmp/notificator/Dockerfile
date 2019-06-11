FROM golang

RUN mkdir -p /go/src/github.com/lexkwan/golang-practice/go-kit-Exmp

ADD . /go/src/github.com/lexkwan/golang-practice/go-kit-Exmp

RUN go get  -t -v ./...
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher -run github.com/lexkwan/golang-practice/go-kit-Exmp/notificator/cmd  -watch github.com/lexkwan/golang-practice/go-kit-Exmp/notificator
