FROM golang:1.22.1-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main .

FROM alpine:3.14
WORKDIR /app
COPY --from=builder /app/main .
COPY /configs /app/configs
CMD ["/app/main"]
