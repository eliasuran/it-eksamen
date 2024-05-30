package main

// TODO: legg til `json:""` for å gjøre dataen til det formatet jeg vil ha i json

type Categories struct {
	Categories []Category
}

type Category struct {
	Name          string
	Link          string
	SubCategories []SubCategory
}

type SubCategory struct {
	Name string
	Link string
}

type Products struct {
	Products []Product
}

type Product struct {
	Id          int
	Title       string
	Category    string
	SubCategory string
	ImageLink   string
	Price       float32
	UnitPrice   float32
	UnitType    string
}
