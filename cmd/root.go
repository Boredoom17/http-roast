package cmd

import (
	"encoding/json"
	"fmt"
	"http-roast/analyzer"
	"http-roast/roaster"
	"os"

	"github.com/spf13/cobra"
)

var jsonOutput bool

var rootCmd = &cobra.Command{
	Use:   "http-roast [url]",
	Short: "Roasts your website's technical quality",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		result, err := analyzer.Analyze(url)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		score := roaster.ScoreResult(result)

		if jsonOutput {
			out := map[string]interface{}{
				"url":           result.URL,
				"status_code":   result.StatusCode,
				"response_time_ms": result.ResponseTime.Milliseconds(),
				"headers":       result.Headers,
				"score":         score.Total,
			}
			b, _ := json.MarshalIndent(out, "", "  ")
			fmt.Println(string(b))
		} else {
			roaster.Roast(result, score)
		}
	},
}

func Execute() {
	rootCmd.PersistentFlags().BoolVar(&jsonOutput, "json", false, "Output as JSON")
	rootCmd.Execute()
}

