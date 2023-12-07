package amazon

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"price-comparison/helper"
	"strings"

	"github.com/gocolly/colly"
)

type Product struct {
	Source    string `json:"source"`
	Link      string `json:"link"`
	Name      string `json:"name"`
	Price     string `json:"price"`
	ImageLink string `json:"imagelink"`
}

func AmazonProduct() {
	c := colly.NewCollector()
	trendyolPath := os.Getenv("TRENDYOL")
	amazonPath := os.Getenv("AMAZON")
	hepsiburadaPath := os.Getenv("HEPSIBURADA")

	fmt.Println("path : ", trendyolPath)

	products := []helper.Product{}

	//-----------------------------------------Hepsiburada--------------------------------------------\\

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")
	})

	c.OnHTML("section.detail-main ", func(h *colly.HTMLElement) {
		split := strings.Split(h.ChildText(".price"), "\n")
		price := strings.TrimSpace(split[0])
		i := helper.Product{
			Source:    "Hepsiburada",
			Link:      hepsiburadaPath,
			Name:      h.ChildText("h1"),
			Price:     price,
			ImageLink: h.ChildAttr("img", "src"),
		}
		oldProduct, _ := helper.GetBySourcePrice("Hepsiburada")
		if oldProduct.Price == price {
			fmt.Println("The price has not changed so it was not saved to csv.")
		} else {
			if oldProduct.Price > price {
				fmt.Println("Price discount")
				products = append(products, i)
				fmt.Println("Web Site : ", i.Source)

			} else if oldProduct.Price < price {
				fmt.Println("Price increase")
				products = append(products, i)
				fmt.Println("Web Site : ", i.Source)

			}

		}
	})

	err := c.Visit(hepsiburadaPath)
	if err != nil {
		log.Fatal(err)
	}

	//-----------------------------------------Amazon--------------------------------------------\\

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")
	})

	c.OnHTML("div.a-container ", func(h *colly.HTMLElement) {
		split := strings.Split(h.ChildText(".a-section"), "TL")
		price := strings.TrimSpace(split[0])
		i := helper.Product{
			Source:    "Amazon",
			Link:      amazonPath,
			Name:      h.ChildText("h1"),
			Price:     price,
			ImageLink: h.ChildAttr("img", "src"),
		}
		oldProduct, _ := helper.GetBySourcePrice("Amazon")
		if oldProduct.Price == price {
			fmt.Println("The price has not changed so it was not saved to csv.")
		} else {
			if oldProduct.Price > price {
				fmt.Println("Price discount")
				products = append(products, i)
				fmt.Println("Web Site : ", i.Source)

			} else if oldProduct.Price < price {
				fmt.Println("Price increase")
				products = append(products, i)
				fmt.Println("Web Site : ", i.Source)

			}

		}

	})

	err = c.Visit(amazonPath)
	if err != nil {
		log.Fatal(err)
	}

	//-----------------------------------------Trendyol--------------------------------------------\\

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")
	})

	c.OnHTML("div.product-detail-container", func(h *colly.HTMLElement) {
		split := strings.Split(h.ChildText("span.prc-dsc"), "TL")
		price := strings.TrimSpace(split[0])
		i := helper.Product{
			Source:    "Trendyol",
			Link:      trendyolPath,
			Name:      h.ChildText("h1"),
			Price:     price,
			ImageLink: h.ChildAttr("img", "src"),
		}
		oldProduct, _ := helper.GetBySourcePrice("Trendyol")
		if oldProduct.Price == price {
			fmt.Println("The price has not changed so it was not saved to csv.")
		} else {
			if oldProduct.Price > price {
				fmt.Println("Price discount")
				products = append(products, i)
				fmt.Println("Web Site : ", i.Source)

			} else if oldProduct.Price < price {
				fmt.Println("Price increase")
				products = append(products, i)
				fmt.Println("Web Site : ", i.Source)

			}

		}

	})

	err = c.Visit(trendyolPath)
	if err != nil {
		log.Fatal(err)
	}

	if len(products) == 0 {
		fmt.Println("Price not changed")
	} else {
		helper.ExcelWriter(products)

	}

	data, err := json.MarshalIndent(products, " ", "")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("All Products")
	fmt.Println(string(data))

}
