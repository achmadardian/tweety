# base build
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .
RUN go build -o myapp .

# final build
FROM alpine:3.21.3

WORKDIR /app
COPY --from=builder /app/myapp .

CMD [ "./myapp" ]  