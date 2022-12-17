FROM golang:1.17

WORKDIR /opt/gzerolearn

ADD  . /opt/gzerolearn

RUN go mod tidy

RUN go build -o greet ./greet.go

EXPOSE 8888

CMD ["/opt/gzerolearn/greet"]