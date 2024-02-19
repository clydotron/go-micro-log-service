# FROM golang:1.21-alpine as builder

# RUN mkdir /app
# COPY . /app
# WORKDIR /app
# RUN CGO_ENABLED=0 go build -o logService ./...
# RUN chmod +x /app/logService

# build as a tiny docker image:
FROM alpine:latest

RUN mkdir /app
COPY logService /app
CMD ["/app/logService"]