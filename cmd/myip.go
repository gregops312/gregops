package cmd

import (
	"encoding/json"
	"fmt"
	output "gregops/pkg"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

const myipCmdName = "myip"

var myipCmd = &cobra.Command{
	Use:   myipCmdName,
	Short: "Get your public IP address",
	Long: fmt.Sprintf(`Get your public IP address by making a request to ipinfo.io.

Examples:
  %s %s
  %s %s --detailed`, CliName, myipCmdName, CliName, myipCmdName),
	Run: func(cmd *cobra.Command, args []string) {
		detailed, _ := cmd.Flags().GetBool("detailed")
		formatStr, _ := cmd.Flags().GetString("output")

		format, err := output.ParseFormat(formatStr)
		if err != nil {
			fmt.Printf("Invalid format: %v\n", err)
			return
		}

		formatter := output.NewWithWriter(format, cmd.OutOrStdout())
		if err := getMyIP(detailed, formatter); err != nil {
			formatter.PrintError(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(myipCmd)
	myipCmd.Flags().BoolP("detailed", "d", false, "Get detailed IP information including location and ISP")
	myipCmd.Flags().StringP("output", "o", "text", "Output format (text, json)")
}

func getMyIP(detailed bool, formatter *output.Formatter) error {
	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Choose endpoint based on detailed flag
	var url string
	if detailed {
		url = "https://ipinfo.io/json"
	} else {
		url = "https://ipinfo.io/ip"
	}

	// Make request to ipinfo.io
	resp, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-200 status code: %d", resp.StatusCode)
	}

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	if detailed {
		return printDetailedInfo(body, formatter)
	}

	// Print just the IP address (trim any whitespace)
	ip := strings.TrimSpace(string(body))
	return formatter.PrintString(ip)
}

func printDetailedInfo(body []byte, formatter *output.Formatter) error {
	var info struct {
		IP       string `json:"ip"`
		City     string `json:"city"`
		Region   string `json:"region"`
		Country  string `json:"country"`
		Location string `json:"loc"`
		Org      string `json:"org"`
		Postal   string `json:"postal"`
		Timezone string `json:"timezone"`
	}

	if err := json.Unmarshal(body, &info); err != nil {
		return fmt.Errorf("failed to parse JSON response: %w", err)
	}

	// Create output data structure
	data := make(map[string]interface{})
	data["IP Address"] = info.IP

	if info.City != "" {
		data["City"] = info.City
	}
	if info.Region != "" {
		data["Region"] = info.Region
	}
	if info.Country != "" {
		data["Country"] = info.Country
	}
	if info.Location != "" {
		data["Location"] = info.Location
	}
	if info.Org != "" {
		data["ISP/Org"] = info.Org
	}
	if info.Postal != "" {
		data["Postal Code"] = info.Postal
	}
	if info.Timezone != "" {
		data["Timezone"] = info.Timezone
	}

	return formatter.PrintKeyValue(data)
}
