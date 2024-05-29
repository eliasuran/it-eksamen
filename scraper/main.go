package main

import (
	"fmt"
	"time"
)

func run() {
	categories := getCategories()

	products := getProducts(categories.categories)

	fmt.Println(products)
}

func main() {
	start := time.Now()
	run()

	fmt.Println("Elapsed:", time.Now().Sub(start))
}
