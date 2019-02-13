# Timing example

It is incredible useful to record the time of
an application and this example does so. Since
one of Golang's selling points is performance 
this is a handy pattern. 

## Usage

```
// Start timer
start := time.Now()

// Operation you wish to time
r := new(big.Int)
fmt.Println(r.Binomial(1000, 10))

// Stop timer
log.Printf("Binomial took %s", time.Since(start))
```

## Get Dependencies

```
go get ./...
```

## Run the code

```
go run main.go
```

### References

[Coderwall measuring execution](https://coderwall.com/p/cp5fya/measuring-execution-time-in-go)