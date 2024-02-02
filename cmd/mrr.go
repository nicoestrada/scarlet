/*
Copyright © 2024 NICO ESTRADA estradanicolas@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/payout"
)

func getMRR() {
	//stored in env
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
	stripe.Key = os.Getenv("STRIPE_API_KEY")

	params := &stripe.PayoutListParams{
		Status: stripe.String("paid"),
	}

	//prevents auto pagination
	params.Single = true

	params.Limit = stripe.Int64(100)

	i := payout.List(params)
	for i.Next() {
		c := i.Payout()
		f := float64(c.Amount) / float64(100)
		fmt.Println("$", f, " - ", c.Status)
	}

}

// mrrCmd represents the mrr command
var mrrCmd = &cobra.Command{
	Use:   "mrr",
	Short: "Get your MRR stats",
	Long:  `This is a working cmd that fetches your most recent 100 payouts and statuses`,
	Run: func(cmd *cobra.Command, args []string) {
		getMRR()
	},
}

func init() {

	rootCmd.AddCommand(mrrCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mrrCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mrrCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
