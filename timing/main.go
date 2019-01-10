package main

import (
	"fmt"
	"log"
	"math/big"
	"time"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %dms", name, elapsed.Nanoseconds()/1000)
}

func factorial(n *big.Int) (result *big.Int) {
	defer timeTrack(time.Now(), "factorial")
	result = big.NewInt(1)
	var one big.Int
	one.SetInt64(1)
	for n.Cmp(&big.Int{}) == 1 {
		result.Mul(result, n)
		n.Sub(n, &one)
	}
	return n
}

func main() {
	start := time.Now()

	r := new(big.Int)
	fmt.Println(r.Binomial(1000, 10))

	log.Printf("Binomial took %s", time.Since(start))

	r2 := new(big.Int)
	fmt.Println(r2.Binomial(1000, 10))

	elapsed2 := time.Since(start)
	log.Printf("Binomial took %dms", elapsed2.Nanoseconds()/1000)

	factorial(big.NewInt(100))
}
