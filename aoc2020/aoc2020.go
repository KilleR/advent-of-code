package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Print("Running ... ")
	start := time.Now()

	day1()

	fmt.Println("done in", time.Since(start))
}
