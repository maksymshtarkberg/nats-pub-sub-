# Dockerfile for publisher
FROM golang:1.22.2

WORKDIR /app
COPY . .

RUN if [ ! -f go.mod ]; then go mod init publisher; fi
RUN go get github.com/nats-io/nats.go
RUN go build -o publisher .

CMD ["./publisher"]
