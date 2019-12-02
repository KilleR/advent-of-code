package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	day1()

	fmt.Println(time.Since(start))
}
