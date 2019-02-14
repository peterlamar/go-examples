# REST API in golang

Simple example of a REST api implemented in golang with the standard library.

## Usage

```golang
func main() {
	http.HandleFunc("/", handler)
    http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
    // function logic
}
```

## Run the code

go run main.go

## Test

curl -H "Content-Type: application/xml" -X GET http://localhost:3000/?name=Alice 

## References

* [Go http library](https://golang.org/pkg/net/http/)
* [Go Writing Web Apps](https://golang.org/doc/articles/wiki/#tmp_4)