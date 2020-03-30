FROM golang:latest

LABEL maintainer="Renish <renishb10@gmail.com>"

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV PORT 7000

RUN go build

CMD ["./foxg-accesslog-service"]