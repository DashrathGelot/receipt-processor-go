FROM golang:1.21.0

WORKDIR /fetch
COPY . .

RUN go mod download

RUN go build -o /fetch/receipt-processor

EXPOSE 8080

CMD ["./receipt-processor"]
