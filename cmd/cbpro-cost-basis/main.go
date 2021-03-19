package main

import (
	"bytes"
	"fmt"
	"os"
	"text/template"

	"github.com/spf13/cobra"
	"github.com/swhite24/cbpro-cost-basis/pkg/config"
	"github.com/swhite24/cbpro-cost-basis/pkg/costbasis"
)

var (
	rootCmd *cobra.Command
)

func init() {
	var startDate, endDate, product string
	var key, passphrase, secret string

	tmpl, _ := template.New("output").Parse(`
Product: {{ .Product }}
Start Date: {{ .Start }}
End Date: {{ .End }}
Total Amount Purchased: {{ .ProductPurchased }}
Total Cost: {{ .TotalCost }}
Cost Basis: {{ .AverageCost }}
`)
	rootCmd = &cobra.Command{
		Use:   "cbpro-cost-basis",
		Short: "cbpro-cost-basis helps calculate cost basis on purchases from Coinbase Pro over a specified period of time",
		Run: func(cmd *cobra.Command, args []string) {
			var output bytes.Buffer

			c, err := config.InitializeConfig(cmd.Flags())
			if err != nil {
				fmt.Println("failed to calculate cost basis")
				fmt.Println(err)
				os.Exit(1)
			}

			info, err := costbasis.Calculate(c)
			if err != nil {
				fmt.Println("failed to calculate cost basis")
				fmt.Println(err)
				os.Exit(1)
			}

			tmpl.Execute(&output, info)
			fmt.Println(output.String())
		},
	}

	rootCmd.Flags().StringVar(&key, "key", "", "Coinbase Pro API key")
	rootCmd.Flags().StringVar(&passphrase, "passphrase", "", "Coinbase Pro API key passphrase")
	rootCmd.Flags().StringVar(&secret, "secret", "", "Coinbase Pro API key secret")

	rootCmd.Flags().StringVar(&startDate, "start", "", "Start date of order fills to calculate cost basis. (2021-01-01)")
	rootCmd.Flags().StringVar(&endDate, "end", "", "End date of order fills to calculate cost basis. (2021-01-01)")
	rootCmd.Flags().StringVar(&product, "product", "BTC-USD", "Product to use when calculating")
}

func main() {
	rootCmd.Execute()
}
