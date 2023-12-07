package main

import (
	"bufio"
	"fmt"
	"os"
	"price-comparison/amazon"
	"strings"
	"time"

	"github.com/go-co-op/gocron"
)

func cronJobs(trendyolPath string, amazonPath string, hepsiburadaPath string) {
	s := gocron.NewScheduler(time.UTC)

	s.Every(60).Seconds().Do(func() {
		amazon.AmazonProduct(trendyolPath, amazonPath, hepsiburadaPath)
	})

	s.StartBlocking()
}

func main() {
	fmt.Print("Write Trendyol product link : ")
	trendyolPath, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	fmt.Print("Write Amazon product link : ")
	amazonPath, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	fmt.Print("Write Hepsiburada product link : ")
	hepsiburadaPath, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	//Clear strings \n
	trendyolPath = strings.TrimSpace(trendyolPath)
	amazonPath = strings.TrimSpace(amazonPath)
	hepsiburadaPath = strings.TrimSpace(hepsiburadaPath)

	cronJobs(trendyolPath, amazonPath, hepsiburadaPath)
}
