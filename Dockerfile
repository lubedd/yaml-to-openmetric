FROM golang:1.17

COPY . /go/src/app

WORKDIR /go/src/app/cmd/openmetric

RUN go build -o openmetric main.go

EXPOSE 8080

CMD ["./openmetric"]