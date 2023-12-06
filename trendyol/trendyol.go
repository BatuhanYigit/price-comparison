package trendyol

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
)

type Item struct {
	Link      string `json:"link"`
	Name      string `json:"name"`
	Price     string `json:"price"`
	ImageLink string `json:"imagelink"`
}

func TrendyolProduct() {
	c := colly.NewCollector()

	items := []Item{}

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")
	})

	c.OnHTML("section.detail-main ", func(h *colly.HTMLElement) {
		split := strings.Split(h.ChildText(".price"), "\n")
		price := strings.TrimSpace(split[0])
		i := Item{
			Link:      h.ChildAttr("a", "href"),
			Name:      h.ChildText("h1"),
			Price:     price,
			ImageLink: h.ChildAttr("img", "src"),
		}
		items = append(items, i)
	})

	err := c.Visit("https://www.hepsiburada.com/steelseries-arctis-1-wireless-oyuncu-kulakligi-usb-c-wireless-ps4-pc-nintendo-switch-android-uyumlu-pm-HB00000NAMDL")
	if err != nil {
		log.Fatal(err)
	}

	data, err := json.MarshalIndent(items, " ", "")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
