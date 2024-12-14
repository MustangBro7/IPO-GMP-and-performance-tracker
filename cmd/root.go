package cmd

import (
	"fmt"
	"os"

	ipo_tracker "ipo_tracker/ipos"

	"github.com/mbndr/figlet4go"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ipo_tracker",
	Short: "Track IPOs and their performance",
	Long: `IPO Tracker is a command-line tool that fetches and displays
IPO-related data, including upcoming IPOs, current GMPs, and historical performance.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Add commands to the root command
	rootCmd.AddCommand(upcomingCmd)
	rootCmd.AddCommand(mainGmpCmd)
	rootCmd.AddCommand(smeGmpCmd)
}

// Command to fetch upcoming IPOs
var upcomingCmd = &cobra.Command{
	Use:   "upcoming",
	Short: "Fetch upcoming IPO data",
	Run: func(cmd *cobra.Command, args []string) {

		ascii := figlet4go.NewAsciiRender()

		// Adding the colors to RenderOptions
		options := figlet4go.NewRenderOptions()
		options.FontColor = []figlet4go.Color{
			// Colors can be given by default ansi color codes...
			figlet4go.ColorGreen,
		}

		renderStr, _ := ascii.RenderOpts("Upcoming", options)
		fmt.Print(renderStr)
		headers, rows := ipo_tracker.Upcoming("https://www.investorgain.com/report/live-ipo-gmp/331/current/", []int{0, 1, 2, 3, 7, 8, 10})
		ipo_tracker.Render(headers, rows)
		renderStr, _ = ascii.RenderOpts("Closed", options)
		fmt.Print(renderStr)
		headers, rows = ipo_tracker.Upcoming("https://www.investorgain.com/report/live-ipo-gmp/331/close/", []int{0, 1, 2, 3, 7, 8, 10})
		ipo_tracker.Render(headers, rows)
	},
}

// Command to fetch main GMP data
var mainGmpCmd = &cobra.Command{
	Use:   "main",
	Short: "Fetch main IPO GMP data",
	Run: func(cmd *cobra.Command, args []string) {
		headers, rows := ipo_tracker.GetGMP("https://www.investorgain.com/report/ipo-performance-history/486/ipo/", []int{0, 5, 6, 8}, "main")
		ipo_tracker.Render(headers, rows)
	},
}

// Command to fetch SME GMP data
var smeGmpCmd = &cobra.Command{
	Use:   "sme",
	Short: "Fetch SME IPO GMP data",
	Run: func(cmd *cobra.Command, args []string) {
		headers, rows := ipo_tracker.GetGMP("https://www.investorgain.com/report/ipo-performance-history/486/sme/", []int{0, 5, 6, 8}, "sme")
		ipo_tracker.Render(headers, rows)
	},
}
