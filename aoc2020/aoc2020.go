package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Running ... ")
	start := time.Now()

	day12()

	fmt.Println("done in", time.Since(start))
}
