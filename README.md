# analytics-service

Service uses Amplitude API

### Host
```
grpc port: 65000
http port: 65001
```

### Usage
**Start app**
    ```
    go run ./cmd/app/main.go
    ```
    
**Pb's generate**
    ```
    make generate-pb
    ```

**Generate swagger docs**
    ```
    make generate-docs
    ```

### Required soft
```
golang 20.*
jaeger
```
**For pb's generate** 
```
protoc-gen-go v2
protoc-gen-go-grpc v2
protoc-gen-grpc-gateway v2
protoc-gen-openapiv2
protoc
make
```
**For generate swagger docs**
```
make
redoc-cli
```
