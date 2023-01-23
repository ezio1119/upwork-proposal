FROM golang:alpine as builder
WORKDIR /build

COPY go.* .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o consumer /build/cmd/consumer/consumer.go
RUN CGO_ENABLED=0 GOOS=linux go build -o producer /build/cmd/producer/producer.go

FROM alpine
WORKDIR /app

COPY --from=builder /build/consumer .
COPY --from=builder /build/producer .
COPY --from=builder /build/cmd/producer/text.csv .