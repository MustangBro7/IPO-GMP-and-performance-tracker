package ipo_tracker

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

// Table styling with Lipgloss
var (
	// Header Style: Neon Gradient with Bold Text
	headerStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#C792EA")).Align(lipgloss.Center).
			Padding(0, 2) // Vibrant purple text

	rowStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF")). // White text for clarity
			Padding(0, 2)                          // Padding for row spacing

	rowAltStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#ABB2BF")). // Light gray for alternate rows
			Padding(0, 2)                          // Padding for row spacing

)

// Function to scrape and display data as a table
func GetGMP(url string, targetColumns []int, reqType string) ([]string, [][]string) {

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to fetch the URL: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error: HTTP Status %d\n", resp.StatusCode)
	}

	// Parse the HTML document with GoQuery
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalf("Failed to parse the HTML document: %v", err)
	}

	var data [][]string
	var headData []string
	var tickerData []string
	// var ipos_now [][]string
	bracketRegex := regexp.MustCompile(`\(([^,)]+)`)
	doc.Find("table tr").Each(func(i int, row *goquery.Selection) {
		var rowData []string
		for _, colIndex := range targetColumns {
			col := row.Find("td").Eq(colIndex)
			if col.Length() > 0 {
				col1 := strings.TrimSpace(col.Text())
				col1 = strings.ReplaceAll(col1, "₹", "")
				col1 = strings.TrimSpace(col1)
				col1 = strings.ReplaceAll(col1, "[email protected]", "")
				// col1 = strings.ReplaceAll(col1, ")", "")
				rowData = append(rowData, strings.TrimSpace(col1))

			} else {
				rowData = append(rowData, "") // Empty if column not found
			}
		}
		if len(rowData) > 0 {
			data = append(data, rowData)
		}
	})
	data = data[2:]
	if len(data) > 16 {
		data = data[:16]
	}
	for i, ticker := range data {
		if len(ticker) == 0 {
			continue // Skip empty rows
		}

		matches := bracketRegex.FindStringSubmatch(ticker[0])
		if len(matches) > 1 {
			colText := matches[1] // Extract only the text inside the first bracket
			tickerData = append(tickerData, colText)
			// fmt.Println(ticker)
			if len(matches) > 1 {
				colText := matches[1] // Extract the text inside the first bracket
				ticker = append(ticker[:0], append([]string{colText}, ticker[0:]...)...)
			} else {
				ticker = append(ticker, "") // Add an empty column for consistency
			}
			data[i] = ticker
		}
	}
	// for _, i := range data {
	// 	fmt.Println(&i)
	// }
	var tickerResult map[string]string
	if reqType == "main" {
		tickerResult = fetchPrices(tickerData)
	} else {
		tickerResult = fetchSMEPrices(tickerData)
	}
	doc.Find("table.table-bordered.table-striped.table-hover.w-auto").Each(func(i int, table *goquery.Selection) { // Replace 'class-name' with your target table's class

		table.Find("thead").Each(func(i int, row *goquery.Selection) {
			for _, colIndex := range targetColumns {
				col := row.Find("th").Eq(colIndex)
				if col.Length() > 0 {
					headData = append(headData, strings.TrimSpace(col.Text()))
				} else {
					headData = append(headData, "") // Empty if column not found
				}
			}
		})
	})
	headData = append(headData[:0], append([]string{"Ticker"}, headData[0:]...)...)
	for _, value := range data {
		value[4] = tickerResult[value[0]]
	}
	headData = append(headData, "Profit/Loss")
	for i, value := range data {
		listPrice, err := strconv.ParseFloat(value[2], 64)
		if err != nil {
			fmt.Println("Error:", err)
		}
		curprice := strings.TrimSpace(value[4])
		currentPrice, err := strconv.ParseFloat(curprice, 64)
		pl := currentPrice - listPrice
		percentage := (pl / listPrice) * 100
		if err != nil {
			fmt.Println("Error:", err)
			pl = 0.0
			percentage = 0.0
		}

		out := fmt.Sprintf("%f (%.2f%%)\n", pl, percentage)
		// println(out)
		value = append(value, out)
		data[i] = value
	}
	headData[4], headData[3] = "Current Price", "Listing Gains"
	return headData, data
}

func Render(headers []string, rows [][]string) {
	// rows = rows[1:] // Skip the first row if needed

	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
		StyleFunc(
			func(row, col int) lipgloss.Style {
				// Header row
				if row == -1 {
					return headerStyle
				}

				// Conditional styling for "SME" in any column
				if strings.Contains(rows[row][col], "SME") {
					return lipgloss.NewStyle().
						Foreground(lipgloss.Color("#FFA500")).Padding(0, 2)
				}
				content := rows[row][col]
				percentage := extractPercentage(content)
				switch {
				case percentage > 0:
					return lipgloss.NewStyle().
						Foreground(lipgloss.Color("46")).Padding(0, 2) // Green for positive
				case percentage == 0:
					if row%2 == 0 {
						return rowStyle
					}
					return rowAltStyle
				case percentage < 0:
					return lipgloss.NewStyle().
						Foreground(lipgloss.Color("196")).Padding(0, 2) // Red for negative
				}

				// Default alternating row styling
				if row%2 == 0 {
					return rowStyle
				}
				return rowAltStyle
			}).
		Headers(headers...).
		Rows(rows...)

	fmt.Println(t)
}

func extractPercentage(content string) int {

	start := strings.LastIndex(content, "(")
	end := strings.LastIndex(content, "%")
	if start != -1 && end != -1 && end > start {
		percentageStr := content[start+1 : end]                    // Extract "14.97" from "(14.97%)"
		percentageStr = strings.ReplaceAll(percentageStr, ",", "") // Handle commas
		percentage, err := strconv.ParseFloat(percentageStr, 64)
		if err == nil {
			return int(percentage)
		}
	}
	return 0
}

func getData(symbol string) string {

	scriptPath := "/Volumes/X10 Pro/projects/ipo_tracker/ipos/price_fetcher.py"
	cmd := exec.Command("/Users/abhinavmohan/myenv/bin/python3", scriptPath, symbol)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error running Python script:", err)
		return err.Error()
	}

	return out.String()
}

func fetchPrices(symbols []string) map[string]string {
	// Using a WaitGroup to wait for all goroutines to complete
	var wg sync.WaitGroup
	results := make(map[string]string)
	mu := &sync.Mutex{} // Mutex to safely write to the shared map

	for _, symbol := range symbols {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			price := getData(s) // Fetch price for the symbol
			mu.Lock()
			results[s] = price
			mu.Unlock()
		}(symbol)
	}

	wg.Wait()
	return results
}

func fetchSMEPrices(symbols []string) map[string]string {
	// Using a WaitGroup to wait for all goroutines to complete
	var wg sync.WaitGroup
	results := make(map[string]string)
	mu := &sync.Mutex{} // Mutex to safely write to the shared map

	for _, symbol := range symbols {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			price := getSMEData(s) // Fetch price for the symbol
			mu.Lock()
			results[s] = price
			mu.Unlock()
		}(symbol)
	}

	wg.Wait()
	return results
}

func Upcoming(url string, targetColumns []int) ([]string, [][]string) {

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to fetch the URL: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error: HTTP Status %d\n", resp.StatusCode)
	}

	// Parse the HTML document with GoQuery
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalf("Failed to parse the HTML document: %v", err)
	}
	var data [][]string
	var headData []string
	doc.Find("table tr").Each(func(i int, row *goquery.Selection) {
		var rowData []string

		for _, colIndex := range targetColumns {
			col := row.Find("td").Eq(colIndex)
			if col.Length() > 0 {
				col1 := strings.ReplaceAll(col.Text(), "[email protected]", "")
				rowData = append(rowData, strings.TrimSpace(col1))

			} else {
				rowData = append(rowData, "") // Empty if column not found
			}
		}
		if len(rowData) > 0 {
			data = append(data, rowData)
		}

	})
	data = data[1:]
	if len(data) > 16 {
		data = data[:16]
	}
	doc.Find("table.table-bordered.table-striped.table-hover.w-auto").Each(func(i int, table *goquery.Selection) {

		table.Find("thead").Each(func(i int, row *goquery.Selection) {
			for _, colIndex := range targetColumns {
				col := row.Find("th").Eq(colIndex)
				if col.Length() > 0 {
					headData = append(headData, strings.TrimSpace(col.Text()))
				} else {
					headData = append(headData, "") // Empty if column not found
				}
			}
		})
	})

	return headData, data
}

func getSMEData(symbol string) string {
	scriptPath := "/Volumes/X10 Pro/projects/ipo_tracker/ipos/sme_fetcher.py"
	cmd := exec.Command("/Users/abhinavmohan/myenv/bin/python3", scriptPath, symbol)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error running Python script:", err)
		return err.Error()
	}

	return out.String()

}
