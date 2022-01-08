#docker build -t skarakasoglu/g-case-challenge:1.0.0
FROM golang:1.17

WORKDIR /go/src/github.com/skarakasoglu/g-case-challenge
ADD . /go/src/github.com/skarakasoglu/g-case-challenge
RUN chmod +x run.sh

CMD ["/bin/bash", "-c", "./run.sh"]