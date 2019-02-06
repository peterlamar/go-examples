## Go with Gin & swagger

Swagger is a method of documenting apis to make them easily accessible and
viewable from a web browser. Swagger and Gin work well together and enable a
straightforward pattern of Swagger docs on Golang. This example highlights this
capability

### Install Go Swagger

Go to `https://github.com/swaggo/gin-swagger` for the documentation.  You need to run `go get -u github.com/swaggo/swag/cmd/swag` to install the swag command.  After this command has been downloaded make sure that your go/bin folder is in your path.

### Generating Swagger Documentation

To generate the swagger docs you can run

```
swag init
```

from the root directory of the project and it should create a docs folder and put the documentation in there.

### Run the code

```
go build && ./ginswagger
```

### Swagger
The swagger page for this app should be served here: http://localhost:8080/swagger/index.html


### Troubleshooting

* Occasionally when developing it is useful to remove the docs directory and executable and regenerate the
swagger docs to show recent changes
* Make sure included structs are exported as swag init doesn't return error messages but silently fails
