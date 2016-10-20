# golang protobuf + grpc gw example

## Setup grpc-gateway

https://github.com/grpc-ecosystem/grpc-gateway

```
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go
```

## Generate gRPC stub

```
$ cd $GOPATH/src/github.com/mrgleam/protobuf-gw-example/pb
$ protoc -I/usr/local/include -I. \
 -I$GOPATH/src \
 -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
 --go_out=Mgoogle/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:. \
 your_service.proto
```

It will generate a stub file `your_service.pb.go` .

```
$ cd $GOPATH/src/github.com/mrgleam/protobuf-gw-example/pb
$ protoc -I/usr/local/include -I. \
 -I$GOPATH/src \
 -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
 --grpc-gateway_out=logtostderr=true:. \
 your_service.proto
```

It will generate a reverse proxy `your_service.pb.gw.go` .

## Run services

```
$ go get github.com/mrgleam/protobuf-gw-example
$ cd $GOPATH/src/github.com/mrgleam/protobuf-gw-example/server
$ go run *.go
```

open new terminal

```
$ cd $GOPATH/src/github.com/mrgleam/protobuf-gw-example/proxy
$ go run main.go
$ curl -X POST http://localhost:8080/v1/example/echo -H "Content-Type: text/plain" -d '{"value": "foo"}'
{"value":"my REST echo"}
```
