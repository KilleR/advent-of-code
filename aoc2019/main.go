package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	day2()

	fmt.Println(time.Since(start))
}
