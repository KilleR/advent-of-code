package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Print("Running ... ")
	start := time.Now()

	day2()

	fmt.Println("done in", time.Since(start))
}
