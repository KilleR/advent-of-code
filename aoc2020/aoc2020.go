package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Running ... ")
	start := time.Now()

	day5()

	fmt.Println("done in", time.Since(start))
}
