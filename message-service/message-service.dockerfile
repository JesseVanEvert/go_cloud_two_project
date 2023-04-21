FROM golang:1.18-alpine as builder
RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o main .

RUN chmod +x /app/main

RUN go get github.com/go-sql-driver/mysql

# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/main /app

CMD [ "/app/main" ]




