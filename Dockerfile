FROM golang:1.14.4-alpine3.12 as builder

ENV GOPATH /go
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY main.go .
RUN go build -o main .

FROM alpine

COPY --from=builder /app/main /app/main

WORKDIR /app
RUN chmod +x main

EXPOSE 8080
ENTRYPOINT ["./main"]
