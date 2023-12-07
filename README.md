# Price Comparison Project

This project is a simple price comparison tool that tracks and compares product prices on different e-commerce websites such as Trendyol, Amazon, and Hepsiburada. It utilizes Go programming language and the Go-Cron library for scheduling.

## Features

- **Scheduled Price Checking**: The program periodically checks and compares prices on specified e-commerce websites.
- **CSV Logging**: Changes in prices are logged to a CSV file for historical tracking.
- **Supported Websites**: Currently supports Trendyol, Amazon, and Hepsiburada.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/price-comparison.git
   
2. Navigate to the project directory:

   ```bash
   cd price-comparison

3. Run the main program:

   ```bash
   go run main.go

## Configuration

Modify the main.go file to input the product links for Trendyol, Amazon, and Hepsiburada.
```bash
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


## Dependencies


Go-Cron: Library used for scheduling tasks.
Colly: Scraping framework used for extracting data from websites.
Usage
Input the product links when prompted.
The program will periodically check prices on the specified websites.
Price changes will be logged to a CSV file.
License
This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

Go: The programming language used for this project.
Colly: The scraping framework used in the project.
Go-Cron: The scheduling library used for periodic tasks.
