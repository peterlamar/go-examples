# Timing example

Choice example from this [blog](https://coderwall.com/p/cp5fya/measuring-execution-time-in-go)

The simple example for timing is

```
func timeTrack(start time.Time, name string) {
    elapsed := time.Since(start)
    log.Printf("%s took %s", name, elapsed)
}
```
