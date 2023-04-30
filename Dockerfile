#FROM alpine:latest
#
#RUN apk --no-cache add ca-certificates
#
#WORKDIR /user-service
#COPY bin/ .
#
#EXPOSE 65000
#EXPOSE 65001
#
#ENTRYPOINT ["./user-service"]


FROM golang:alpine as builder

RUN apk add git
ADD . /go/src/user-service
WORKDIR /go/src/user-service
RUN go mod download

COPY . ./


RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /user-service ./cmd/app/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /user-service ./user-service
RUN mkdir ./config
COPY config/public_config.yaml ./config

EXPOSE 65001

ENTRYPOINT ["./user-service"]