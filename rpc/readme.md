# TCP RPC example

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
