/*
Copyright © 2024 NICO ESTRADA estradanicolas@gmail.com
*/
package cmd

import (
	"fmt"
	"math"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/payout"
)

// iterates every payout to a running sum / 12
func getMRR() {
	//stored in env
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
	stripe.Key = os.Getenv("STRIPE_API_KEY")
	params := &stripe.PayoutListParams{
		Status: stripe.String("paid"),
	}
	params.Single = true //prevents auto pagination
	params.Limit = stripe.Int64(100)
	i := payout.List(params)
	total := float64(0)
	for i.Next() {
		c := i.Payout()
		f := float64(c.Amount) / float64(100)
		total += f
	}
	fmt.Println(
		"\n\n Total payouts = $",
		math.Floor(((total * 100) / 100)),
		"\n\n MRR = $",
		math.Floor(((total*100)/100)/12),
	)
}

// mrrCmd represents the mrr command
var mrrCmd = &cobra.Command{
	Use:   "mrr",
	Short: "Get your real MRR",
	Long:  `This is a working cmd that fetches your most recent 100 payouts and statuses`,
	Run: func(cmd *cobra.Command, args []string) {
		getMRR()
	},
}

func init() {
	rootCmd.AddCommand(mrrCmd)
}
