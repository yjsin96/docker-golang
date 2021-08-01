FROM golang:1.16-alpine AS builder
WORKDIR /build
COPY go.mod go.sum ./
COPY . .
RUN GOOS=linux go build -o test-golang

FROM alpine:latest
WORKDIR /app
COPY --from=builder /build/test-golang ./
CMD [ "./test-golang" ]
