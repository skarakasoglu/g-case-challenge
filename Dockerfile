FROM golang:1.17

WORKDIR /go/src/github.com/skarakasoglu/g-case-challenge
ADD . /go/src/github.com/skarakasoglu/g-case-challenge

CMD ["/bin/bash", "-c", "./run.sh"]