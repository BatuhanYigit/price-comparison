package helper

import (
	"encoding/csv"
	"log"
	"os"
)

type Product struct {
	Source    string `json:"source"`
	Link      string `json:"link"`
	Name      string `json:"name"`
	Price     string `json:"price"`
	ImageLink string `json:"imagelink"`
}

func ExcelWriter(products []Product) {
	file, err := os.Create("./products.csv")
	if err != nil {
		log.Fatalln("Failed create csv ", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	headers := []string{
		"Source",
		"Link",
		"Name",
		"Price",
		"ImageLink",
	}
	writer.Write(headers)

	for _, product := range products {
		record := []string{
			product.Source,
			product.Link,
			product.Name,
			product.Price,
			product.ImageLink,
		}

		writer.Write(record)

	}

	defer writer.Flush()
}
