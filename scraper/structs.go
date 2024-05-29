package main

type categories struct {
	categories []category
}

type category struct {
	name string
	link string
}

type products struct {
	products []product
}

type product struct {
	id        int
	title     string
	category  string
	imageLink string
	price     price
}

type price struct {
	price      float32
	unitprice  float32
	unittype   string
	product_id int
}
