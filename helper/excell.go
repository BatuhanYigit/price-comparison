package helper

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
)

type Product struct {
	Source    string `json:"source"`
	Link      string `json:"link"`
	Name      string `json:"name"`
	Price     string `json:"price"`
	ImageLink string `json:"imagelink"`
}

func OpenCreateCsv(filePath string, headers []string) (*os.File, *csv.Writer, error) {
	var file *os.File
	var writer *csv.Writer

	if _, err := os.Stat(filePath); os.IsNotExist(err) {

		file, err = os.Create(filePath)
		if err != nil {
			return nil, nil, err
		}
		writer = csv.NewWriter(file)
		writer.Write(headers)
	} else {

		file, err = os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			return nil, nil, err
		}
		writer = csv.NewWriter(file)
	}

	return file, writer, nil
}

func ExcelWriter(products []Product) {
	filePath := "./products.csv"
	headers := []string{
		"Source",
		"Link",
		"Name",
		"Price",
		"ImageLink",
	}
	file, writer, err := OpenCreateCsv(filePath, headers)
	if err != nil {
		fmt.Println("Error open csv ", err)
	}
	defer file.Close()

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

	writer.Flush()

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

func GetBySource(source string) []Product {
	products, err := ReadCSV()
	if err != nil {
		fmt.Println("CSV okuma hatası:", err)
	}
	var filterProducts []Product

	for _, product := range products {
		if product.Source == source {
			filterProducts = append(filterProducts, product)
		}
	}

	return filterProducts
}

func GetBySourcePrice(source string) (Product, error) {
	products, err := ReadCSV()
	if err != nil {
		fmt.Println("Read error", err)
	}
	var foundProduct Product

	for _, product := range products {
		if product.Source == source {
			foundProduct = product
		}
	}

	if foundProduct == (Product{}) {
		return Product{}, errors.New("Product not found")
	}

	return foundProduct, nil
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
