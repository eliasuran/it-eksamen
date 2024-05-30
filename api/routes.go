package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func httpError(w http.ResponseWriter, err error, message string) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, message, err)
}

func routes(
	mux *http.ServeMux,
	client *sql.DB,
) {
	mux.HandleFunc("GET /products", func(w http.ResponseWriter, r *http.Request) {
		rows, err := client.Query("SELECT * FROM products")
		if err != nil {
			httpError(w, err, "Could not get products from database: %v\n")
			return
		}
		defer rows.Close()

		var products []Product

		for rows.Next() {
			var product Product
			if err := rows.Scan(&product.Id, &product.Title, &product.Category, &product.Subcategory, &product.Imagelink); err != nil {
				httpError(w, err, "Could not read data from database: %v\n")
				return
			}
			products = append(products, product)
		}
		if err = rows.Err(); err != nil {
			httpError(w, err, "Error occured during iteration: %v\n")
			return
		}

		json, err := json.MarshalIndent(products, "", "    ")
		if err != nil {
			httpError(w, err, "Error occured during iteration: %v\n")
			return
		}

		w.WriteHeader(200)
		fmt.Fprintln(w, string(json))
	})

	// mux.HandleFunc("GET /products/{id}", func(w http.ResponseWriter, r *http.Request) {})

	mux.Handle("GET /", http.NotFoundHandler())
}
