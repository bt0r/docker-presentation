FROM golang:1.10-alpine3.8
LABEL maintainer="Thibaut.B <thibaut.bayer@evaneos.com>"


RUN mkdir /var/www
WORKDIR /var/www

RUN apk add git
RUN go get -u github.com/go-pg/pg

COPY src/ .

EXPOSE 80

CMD go run main.go
