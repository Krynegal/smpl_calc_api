FROM golang:1.18-alpine

WORKDIR /calcApp

COPY . .

RUN go mod download
RUN go build -o calcApp ./cmd/server/main.go

EXPOSE 8080

CMD ["./calcApp"]