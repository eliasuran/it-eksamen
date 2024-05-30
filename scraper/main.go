package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func run() {
	categories := getCategories()

	json, _ := json.MarshalIndent(categories, "", "    ")
	fmt.Println(string(json))
}

func main() {
	start := time.Now()
	run()

	fmt.Println("Elapsed:", time.Now().Sub(start))
}
