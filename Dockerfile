FROM golang:1.22.1-alpine AS builder
WORKDIR /app
COPY . .
RUN GOOS=linux go build -o main .

FROM alpine:3.19 as runner
COPY --from=builder /app/main /app/main
ADD /configs /configs 
CMD ["/app/main"]