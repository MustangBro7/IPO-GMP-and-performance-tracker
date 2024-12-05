package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

// Table styling with Lipgloss
var (
	// Header Style: Neon Gradient with Bold Text
	headerStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#C792EA")). // Vibrant purple text
		// Background(lipgloss.Color("#1E1E2F")).       // Dark purple/black background
		Border(lipgloss.ThickBorder()) // Thick border for a strong header
		// BorderForeground(lipgloss.Color("#A1EFD3")). // Light cyan border
		// Padding(0, 2) // Padding for a modern feel

	// Row Style: Subtle Gradient Effect
	rowStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF")). // White text for clarity
		// Background(lipgloss.Color("#282A36")). // Dark background
		Padding(0, 2) // Padding for row spacing

	// Row Alternating Style: Gradient for Rows
	rowAltStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#ABB2BF")). // Light gray for alternate rows
		// Background(lipgloss.Color("#232530")). // Slightly different background
		Padding(0, 2) // Padding for row spacing

	// Border Style: Glowing Rounded Borders
	borderStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#FF00FF")). // Neon magenta border
			Padding(1, 2).                               // Padding inside the table
			Background(lipgloss.Color("#000000"))        // Black background for modern effect

		// Separator: Bright Separator Between Columns

)

// Function to scrape and display data as a table
func gmp(url string, targetColumns []int) ([]string, [][]string) {
	// URL of the website to search
	// Replace with the website URL

	// Fetch the HTML document
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

	// Headers for the table
	// headers := []string{"Name", "Price", "GMP", "Listing Gains", "Open", "Close", "Listing"} // Replace with your desired headers
	// fmt.Print(headers)
	// Data rows extracted from the table
	var data [][]string
	var headData []string
	// var ipos_now [][]string
	doc.Find("table tr").Each(func(i int, row *goquery.Selection) {
		var rowData []string

		// targetColumns := []int{0, 1, 2, 3, 7, 8, 9} // Indices for columns 1, 3, 5 (0-based)
		for _, colIndex := range targetColumns {
			col := row.Find("td").Eq(colIndex)
			if col.Length() > 0 {
				// col.Text() = strings.ReplaceAll(col.Text(), "L@", "")
				col1 := strings.ReplaceAll(col.Text(), "[email protected]", "")
				rowData = append(rowData, strings.TrimSpace(col1))

			} else {
				rowData = append(rowData, "") // Empty if column not found
			}
		}
		if len(rowData) > 0 {
			data = append(data, rowData)
		}

		// fmt.Print(rowData)
		// fmt.Print("\n")
	})
	data = data[1:]
	if len(data) > 16 {
		data = data[:16]
	}
	doc.Find("table.table-bordered.table-striped.table-hover.w-auto").Each(func(i int, table *goquery.Selection) { // Replace 'class-name' with your target table's class

		table.Find("thead").Each(func(i int, row *goquery.Selection) {
			// targetColumns := []int{0, 1, 2, 3, 7, 8, 9} // Indices for columns (0-based)
			for _, colIndex := range targetColumns {
				col := row.Find("th").Eq(colIndex)
				if col.Length() > 0 {
					headData = append(headData, strings.TrimSpace(col.Text()))
				} else {
					headData = append(headData, "") // Empty if column not found
				}
			}

			// fmt.Print(headData)
			// fmt.Print("\n")
		})
	})

	return headData, data
}

func main() {
	headers, rows := gmp("https://www.investorgain.com/report/live-ipo-gmp/331/current/", []int{0, 1, 2, 3, 7, 8, 9})
	renderrrr(headers, rows)
	headers, rows = gmp("https://www.investorgain.com/report/ipo-performance-history/486/all/", []int{0, 5, 6, 8})
	headers[3], headers[2] = "Current Price", "Listing Gains"
	renderrrr(headers, rows)

}

func renderrrr(headers []string, rows [][]string) {
	// rows = rows[1:] // Skip the first row if needed

	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
		StyleFunc(
			func(row, col int) lipgloss.Style {
				// Header row
				if row == -1 {
					return rowStyle
				}

				// Conditional styling for "SME" in any column
				if strings.Contains(rows[row][col], "SME") {
					return lipgloss.NewStyle().
						Foreground(lipgloss.Color("#FFA500")).Padding(0, 2)
				}

				// Check the 4th column (index 3) for percentage-based styling
				if (col == 2 || col == 3) && row < len(rows) {
					// Extract the numeric value from the percentage
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

// Helper function to extract percentage from the 4th column

func extractPercentage(content string) int {
	// Example: "169 (14.97%)"
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
	return 0 // Default to 0 if no valid percentage is found
}

// package main

// import (
// 	"encoding/hex"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"strconv"
// 	"strings"

// 	"github.com/PuerkitoBio/goquery"
// 	"github.com/charmbracelet/lipgloss"
// 	"github.com/charmbracelet/lipgloss/table"
// )

// // Table styling with Lipgloss (same as in your code)
// var (
// 	// Header Style: Neon Gradient with Bold Text
// 	headerStyle = lipgloss.NewStyle().
// 			Bold(true).
// 			Foreground(lipgloss.Color("#C792EA")). // Vibrant purple text
// 		// Background(lipgloss.Color("#1E1E2F")).       // Dark purple/black background
// 		Border(lipgloss.ThickBorder()) // Thick border for a strong header
// 		// BorderForeground(lipgloss.Color("#A1EFD3")). // Light cyan border
// 		// Padding(0, 2) // Padding for a modern feel

// 	// Row Style: Subtle Gradient Effect
// 	rowStyle = lipgloss.NewStyle().
// 			Foreground(lipgloss.Color("#FFFFFF")). // White text for clarity
// 		// Background(lipgloss.Color("#282A36")). // Dark background
// 		Padding(0, 2) // Padding for row spacing

// 	// Row Alternating Style: Gradient for Rows
// 	rowAltStyle = lipgloss.NewStyle().
// 			Foreground(lipgloss.Color("#ABB2BF")). // Light gray for alternate rows
// 		// Background(lipgloss.Color("#232530")). // Slightly different background
// 		Padding(0, 2) // Padding for row spacing

// 	// Border Style: Glowing Rounded Borders
// 	borderStyle = lipgloss.NewStyle().
// 			Border(lipgloss.RoundedBorder()).
// 			BorderForeground(lipgloss.Color("#FF00FF")). // Neon magenta border
// 			Padding(1, 2).                               // Padding inside the table
// 			Background(lipgloss.Color("#000000"))        // Black background for modern effect

// 		// Separator: Bright Separator Between Columns

// )

// func decodeCfEmail(encoded string) string {
// 	data, err := hex.DecodeString(encoded)
// 	if err != nil {
// 		log.Fatalf("Failed to decode hex string: %v", err)
// 	}

// 	key := data[0] // The first byte is the key
// 	decoded := make([]byte, len(data)-1)

// 	for i := 1; i < len(data); i++ {
// 		decoded[i-1] = data[i] ^ key // XOR each byte with the key
// 	}

// 	return string(decoded)
// }

// // Function to scrape and extract headers and rows from the table
// func gmp(url string, targetColumns []int) ([]string, [][]string) {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		log.Fatalf("Failed to fetch the URL: %v", err)
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		log.Fatalf("Error: HTTP Status %d\n", resp.StatusCode)
// 	}

// 	doc, err := goquery.NewDocumentFromReader(resp.Body)
// 	if err != nil {
// 		log.Fatalf("Failed to parse the HTML document: %v", err)
// 	}

// 	var data [][]string
// 	var headData []string

// 	// Extract table headers
// 	doc.Find("table.table-bordered.table-striped.table-hover.w-auto thead").Each(func(i int, row *goquery.Selection) {
// 		for _, colIndex := range targetColumns {
// 			col := row.Find("th").Eq(colIndex)
// 			if col.Length() > 0 {
// 				headData = append(headData, strings.TrimSpace(col.Text()))
// 			} else {
// 				headData = append(headData, "")
// 			}
// 		}
// 	})

// 	// Extract table rows
// 	doc.Find("table.table-bordered.table-striped.table-hover.w-auto tbody tr").Each(func(i int, row *goquery.Selection) {
// 		var rowData []string
// 		for _, colIndex := range targetColumns {
// 			col := row.Find("td").Eq(colIndex)

// 			// Handle Cloudflare obfuscated content
// 			emailLink := col.Find("a.__cf_email__")
// 			if emailLink.Length() > 0 {
// 				cfEmail := emailLink.AttrOr("data-cfemail", "")
// 				if cfEmail != "" {
// 					decodedValue := decodeCfEmail(cfEmail)
// 					decodedValue = strings.ReplaceAll(decodedValue, "L@", "") // Remove "L@" prefix
// 					rowData = append(rowData, decodedValue)
// 				} else {
// 					rowData = append(rowData, "[email protected]")
// 				}
// 			} else {
// 				// Clean the text and extract percentages
// 				content := strings.TrimSpace(col.Text())
// 				if colIndex == 3 { // Assuming 4th column contains the percentage
// 					content = extractPercentageWithValue(content) // Clean and keep percentage
// 				}
// 				rowData = append(rowData, strings.ReplaceAll(content, "L@", "")) // Remove "L@" prefix
// 			}
// 		}

// 		if len(rowData) > 0 {
// 			data = append(data, rowData)
// 		} else {
// 			fmt.Printf("Warning: Empty row detected: %v\n", row.Text()) // Debug output for empty rows
// 		}
// 	})

// 	// Limit to the latest 15 entries
// 	if len(data) > 15 {
// 		data = data[len(data)-15:]
// 	}

// 	return headData, data
// }

// // Render table
// func renderrrr(headers []string, rows [][]string) {
// 	t := table.New().
// 		Border(lipgloss.NormalBorder()).
// 		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
// 		StyleFunc(
// 			func(row, col int) lipgloss.Style {
// 				if row == -1 {
// 					return rowStyle // Header styling
// 				}
// 				if col == 3 && row < len(rows) { // Highlight percentages
// 					percentage := extractPercentage(rows[row][col])
// 					if percentage > 0 {
// 						return lipgloss.NewStyle().Foreground(lipgloss.Color("46")) // Green
// 					} else if percentage < 0 {
// 						return lipgloss.NewStyle().Foreground(lipgloss.Color("196")) // Red
// 					} else {
// 						return lipgloss.NewStyle().Foreground(lipgloss.Color("244")) // Gray
// 					}
// 				}
// 				if row%2 == 0 {
// 					return rowStyle
// 				}
// 				return rowAltStyle
// 			}).
// 		Headers(headers...).
// 		Rows(rows...)

// 	fmt.Println(t)
// }

// // Helper function to extract percentage as an integer
// // Helper function to extract and retain percentage in content
// func extractPercentageWithValue(content string) string {
// 	// Example: "198 (83.33%)"
// 	start := strings.LastIndex(content, "(")
// 	end := strings.LastIndex(content, "%")
// 	if start != -1 && end != -1 && end > start {
// 		percentage := content[start : end+1]        // Extract "(value%)"
// 		value := strings.TrimSpace(content[:start]) // Extract value before the percentage
// 		return value + " " + percentage
// 	}
// 	return content // Return content as-is if no percentage is found
// }

// // Helper function to extract percentage as an integer
// func extractPercentage(content string) int {
// 	start := strings.LastIndex(content, "(")
// 	end := strings.LastIndex(content, "%")
// 	if start != -1 && end != -1 && end > start {
// 		percentageStr := content[start+1 : end]                    // Extract "83.33" from "(83.33%)"
// 		percentageStr = strings.ReplaceAll(percentageStr, ",", "") // Handle commas
// 		percentage, err := strconv.ParseFloat(percentageStr, 64)
// 		if err == nil {
// 			return int(percentage)
// 		}
// 	}
// 	fmt.Printf("Warning: Could not extract percentage from content: %s\n", content) // Debug output
// 	return 0                                                                        // Default to 0 if no valid percentage is found
// }

// func main() {
// 	headers, rows := gmp("https://www.investorgain.com/report/ipo-performance-history/486/all/", []int{0, 4, 5, 6})
// 	renderrrr(headers, rows)
// }
