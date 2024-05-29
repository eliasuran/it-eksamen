package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

func getCategories() categories {
	categories := categories{}

	c := colly.NewCollector()

	// i navigasjonsbaren som kategorier hentes fra, er det først noen elementer som ikke er kategorier
	// bruker dette for å tracke om kategorier har begynt å komme
	categoriesStarted := false

	c.OnHTML("div.sc-adf1bc0c-8", func(e *colly.HTMLElement) {
		href := e.ChildAttr("div > a", "href")
		link := fmt.Sprintf("%s%s", "https://wolt.com", href)

		name := e.ChildText("div.sc-dad41f1f-6")

		// frukt & grønnsaker er den første kategorien, så da kan jeg begynne å legge til kategoriene i instansen av categories
		if name == "Frukt & Grønnsaker" {
			categoriesStarted = true
		}

		if categoriesStarted {
			categories.categories = append(categories.categories, category{name, link})
		}
	})

	c.Visit("https://wolt.com/nb/nor/oslo/venue/wolt-market-sentrum")

	return categories
}

func getProducts(categories []category) products {
	products := products{}

	id := 1

	for _, category := range categories {
		c := colly.NewCollector()

		c.OnHTML("div.sc-32c83f74-3", func(e *colly.HTMLElement) {
			title := e.ChildText("h3.sc-32c83f74-10")
			fmt.Println("Henter data for:", title)
			imageLink := e.ChildAttr("img", "src")

			priceStr := e.ChildText("span.sc-ceacab0-2")
			if priceStr == "" {
				priceStr = e.ChildText("span.sc-ceacab0-0")
			}
			priceFloat, err := strconv.ParseFloat(strings.Split(priceStr, "&")[0], 32)
			if err != nil {
				priceFloat = 0.00
			}
			fmt.Println(priceStr)

			unitPrice := strings.Split(e.ChildText("span.sc-ca7057be-2"), "/")
			var unitPriceFloat float64
			var unitType string

			if len(unitPrice) >= 2 {
				unitPriceFloat, err = strconv.ParseFloat(strings.Split(unitPrice[0], "&")[0], 32)
				if err != nil {
					unitPriceFloat = 0.00
				}
				unitType = unitPrice[1]
			}

			products.products = append(products.products, product{
				id:    id,
				title: title,
				price: price{
					price:      float32(priceFloat),
					unitprice:  float32(unitPriceFloat),
					unittype:   unitType,
					product_id: id,
				},
				category:  category.name,
				imageLink: imageLink,
			})

			id++
		})

		c.Visit(category.link)
	}

	return products
}
