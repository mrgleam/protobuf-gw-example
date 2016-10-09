# golang protobuf + grpc gw example

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
