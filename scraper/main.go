package main

import (
	"fmt"
	"time"
)

func scraper() {
	fmt.Println("hello world")
}

func main() {
	start := time.Now()
	scraper()

	fmt.Println("Elapsed:", time.Now().Sub(start))
}
