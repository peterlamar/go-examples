# Timing example

Choice example from this [blog](https://coderwall.com/p/cp5fya/measuring-execution-time-in-go)

The simple example for timing is

```
func main() {
    start := time.Now()

    r := new(big.Int)
    fmt.Println(r.Binomial(1000, 10))

    elapsed := time.Since(start)
    log.Printf("Binomial took %s", elapsed)
}
```
