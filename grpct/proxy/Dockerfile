FROM golang:1.19-alpine as builder

WORKDIR /app

COPY . .

COPY .env ./

RUN go mod tidy

RUN go build -o main ./cmd

FROM alpine:latest

COPY --from=builder /app/main /main

EXPOSE 8080

CMD ["/main"]
