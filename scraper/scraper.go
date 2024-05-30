package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

func getCategories() Categories {
	categories := Categories{}

	c := colly.NewCollector()

	// i navigasjonsbaren som kategorier hentes fra, er det først noen elementer som ikke er kategorier
	// bruker dette for å tracke om kategorier har begynt å komme
	categoriesStarted := false

	c.OnHTML("div.sc-adf1bc0c-8", func(e *colly.HTMLElement) {
		name := e.ChildText("div.sc-dad41f1f-6")
		name = strings.ReplaceAll(name, "\u0026", "&")

		href := e.ChildAttr("div > a", "href")
		link := fmt.Sprintf("%s%s", "https://wolt.com", href)

		// frukt & grønnsaker er den første kategorien, så da settes categoriesStarted til true
		if name == "Frukt & Grønnsaker" {
			categoriesStarted = true
		}

		// om categoriesStarted er true, hentes underkategorier og dataen legges til i instansen av categories structet
		if categoriesStarted {
			subCategories := getSubCategories(link)
			categories.Categories = append(categories.Categories, Category{name, link, subCategories})
		}
	})

	c.Visit("https://wolt.com/nb/nor/oslo/venue/wolt-market-sentrum")

	return categories
}

func getSubCategories(link string) []SubCategory {
	subCategories := []SubCategory{}

	c := colly.NewCollector()

	c.OnHTML("div.sc-adf1bc0c-8.kWmoAY", func(e *colly.HTMLElement) {
		name := e.ChildText("a > div.sc-dad41f1f-6")
		name = strings.ReplaceAll(name, "\u0026", "&")

		if name != "Alle varer" && !strings.Contains(name, "Kampanjer") {
			href := e.ChildAttr("a", "href")
			link := fmt.Sprintf("%s%s", "https://wolt.com", href)

			subCategories = append(subCategories, SubCategory{name, link})
		}

	})

	c.Visit(link)

	return subCategories
}

func getProducts(products *Products, category Category, subCategory SubCategory, id *int) {
	c := colly.NewCollector()

	c.OnHTML("div.sc-32c83f74-3", func(e *colly.HTMLElement) {
		title := e.ChildText("h3.sc-32c83f74-10")
		fmt.Println("Henter data for:", title)

		imageLink := e.ChildAttr("img", "src")

		priceStr := e.ChildText("span.sc-ceacab0-1")
		priceFloat, err := strconv.ParseFloat(strings.Split(priceStr, "&")[0], 32)
		if err != nil {
			priceFloat = 0
		}

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

		products.Products = append(products.Products, Product{
			Id:          *id,
			Title:       title,
			Price:       float32(priceFloat),
			UnitPrice:   float32(unitPriceFloat),
			UnitType:    unitType,
			Category:    category.Name,
			SubCategory: subCategory.Name,
			ImageLink:   imageLink,
		})

		(*id)++
	})

	c.Visit(subCategory.Link)
}
