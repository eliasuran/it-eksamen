package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func run() {
	categories := getCategories()

	products := &Products{}
	id := 1
	for _, category := range categories.Categories {
		for _, subCategory := range category.SubCategories {
			getProducts(products, category, subCategory, &id)
		}
	}

	db, err := sql.Open("postgres", os.Getenv("NEON_URL"))
	if err != nil {
		log.Fatalf("Error creating postgres client: %v\n", err)
	}

	var wg sync.WaitGroup
	sem := make(chan struct{}, 4)

	for _, product := range products.Products {
		wg.Add(1)
		sem <- struct{}{}

		go func(product Product) {
			defer wg.Done()
			defer func() { <-sem }()
			insertData(db, product)
		}(product)
	}
}

func main() {
	godotenv.Load()

	start := time.Now()
	run()

	fmt.Println("Elapsed:", time.Now().Sub(start))
}
