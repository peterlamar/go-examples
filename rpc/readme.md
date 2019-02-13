# TCP RPC example

This is an example of an RPC (Remote Procedure Call)
in Golang. This is a method of writing a function in one place and calling it in another. Typically it is referred to as a less often used alternative to REST. 

The default transmission protocol for Go RPC is a Go binary format. JSON can be selected as well for a transition method but it will be slower as JSON is more verbose. 

For modern applications, [grpc](https://grpc.io/) is 
typically used in production as the formats are small
and efficient. Regular RPC can be used in development
as grpc requires protobuf definition files which add
complexity to the application. 

## Get Dependencies

This step downloads dependencies needed. The pattern "./..." means start in the
current directory ("./") and find all packages below that directory ("...")

More [info](https://golang.org/doc/articles/go_command.html#tmp_3)

```
go get ./...
```

## Run the code

The [go command](https://golang.org/cmd/go/) allows you to run the go files
directly. Sometimes this is useful when starting a project or running simple
demos that consist of a single file.

1. Start server

```
go run server.go
```

2. Open a new tab in the terminal and start the client

```
go run client.go
```

3. Start typing in the client tab and observe the transferred data

## Reference 

[Go rpc Package](https://golang.org/pkg/net/rpc/)