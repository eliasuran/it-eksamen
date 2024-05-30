package main

import (
	"database/sql"
	"fmt"
)

func insertData(db *sql.DB, product Product) {
	fmt.Println("Legger inn data for:", product.Title)

	query, err := db.Prepare(`
		INSERT INTO products (id, title, category, subcategory, imagelink)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (id)
		DO UPDATE SET
			title = EXCLUDED.title,
			category = EXCLUDED.category,
			subcategory = EXCLUDED.subcategory,
			imagelink = EXCLUDED.imagelink
	`)
	if err != nil {
		fmt.Printf("Error preparing query: %v\n", err)
		return
	}
	defer query.Close()

	_, err = query.Exec(product.Id, product.Title, product.Category, product.SubCategory, product.ImageLink)
	if err != nil {
		fmt.Printf("Error running query: %v\n", err)
		return
	}

	query, err = db.Prepare(`
		INSERT INTO prices (price, unitprice, unittype, product_id)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (product_id)
		DO UPDATE SET
			price = EXCLUDED.price,
			unitprice = EXCLUDED.unitprice,
			unittype = EXCLUDED.unittyp
	`)
	if err != nil {
		fmt.Printf("Error preparing query: %v\n", err)
		return
	}
	defer query.Close()

	_, err = query.Exec(product.Price.Price, product.Price.UnitPrice, product.Price.UnitType, product.Id)
	if err != nil {
		fmt.Printf("Error running query: %v\n", err)
		return
	}
}
