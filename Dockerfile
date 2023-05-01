FROM golang:alpine as builder

RUN apk add git
ADD . /go/src/analytics-service
WORKDIR /go/src/analytics-service
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /analytics-service ./cmd/app/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /analytics-service ./analytics-service
RUN mkdir ./config
COPY config/private_config.yaml ./config

EXPOSE 65001

ENTRYPOINT ["./analytics-service"]