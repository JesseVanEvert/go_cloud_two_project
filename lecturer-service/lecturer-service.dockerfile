# base go image
FROM golang:1.18-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o lecturerService .

RUN chmod +x /app/lecturerService

# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/lecturerService /app

CMD [ "/app/lecturerService" ]