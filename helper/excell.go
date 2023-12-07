package helper

import (
	"encoding/csv"
	"fmt"
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

func GetAll() {
	file, err := os.Open("./products.csv")
	if err != nil {
		fmt.Println("File open error", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	lines, err := reader.ReadAll()

	if err != nil {
		fmt.Println("File read error ", err)
		return
	}

	for _, line := range lines {
		fmt.Println(line)
	}

}

func GetBySource(products []Product, source string) []Product {
	var filterProducts []Product

	for _, product := range products {
		if product.Source == source {
			filterProducts = append(filterProducts, product)
		}
	}

	return filterProducts
}

func ReadCSV() ([]Product, error) {
	filePath := "./products.csv"
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	var products []Product
	headers, err := reader.Read()
	if err != nil {
		return nil, err
	}

	for {
		line, err := reader.Read()
		if err != nil {
			break
		}

		product := Product{}
		for i, value := range line {
			switch headers[i] {
			case "Source":
				product.Source = value
			case "Link":
				product.Link = value
			case "Name":
				product.Name = value
			case "Price":
				product.Price = value
			case "ImageLink":
				product.ImageLink = value
			}
		}
		products = append(products, product)
	}

	return products, nil
}
