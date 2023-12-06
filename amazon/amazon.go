package amazon

import (
	"encoding/json"
	"fmt"
	"log"
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
			Link:      h.ChildAttr("a", "href"),
			Name:      h.ChildText("h1"),
			Price:     price,
			ImageLink: h.ChildAttr("img", "src"),
		}
		products = append(products, i)
	})

	err := c.Visit("https://www.hepsiburada.com/steelseries-arctis-1-wireless-oyuncu-kulakligi-usb-c-wireless-ps4-pc-nintendo-switch-android-uyumlu-pm-HB00000NAMDL")
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
			Link:      h.ChildAttr("a", "href"),
			Name:      h.ChildText("h1"),
			Price:     price,
			ImageLink: h.ChildAttr("img", "src"),
		}
		products = append(products, i)
	})

	err = c.Visit("https://www.amazon.com.tr/SteelSeries-Arctis-Wireless-Gaming-Kulakl%C4%B1k/dp/B09GW7NHL4/ref=sr_1_1?crid=13YG3B9ISWB9N&keywords=steelseries+arctis+7&qid=1701335813&sprefix=steel+%2Caps%2C161&sr=8-1")
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
			Link:      h.ChildAttr("a", "href"),
			Name:      h.ChildText("h1"),
			Price:     price,
			ImageLink: h.ChildAttr("img", "src"),
		}
		products = append(products, i)
	})

	err = c.Visit("https://www.trendyol.com/steelseries/arctis-1-wireless-oyuncu-kulaklik-pc-ps4-ps5-xbox-nintendo-switch-ve-android-p-31622422?boutiqueId=61&merchantId=220053")
	if err != nil {
		log.Fatal(err)
	}

	helper.ExcelWriter(products)

	data, err := json.MarshalIndent(products, " ", "")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("All Products")
	fmt.Println(string(data))

}
