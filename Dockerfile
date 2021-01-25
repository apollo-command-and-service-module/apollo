# Build container
FROM golang:alpine as builder
RUN apk add build-base
WORKDIR /go/src/app
COPY . .
RUN go build -o sm ./service-module/service-module.go

EXPOSE 5000
CMD ["./sm"]