/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
)

func getCustomers() {
	//stored in env
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
	stripe.Key = os.Getenv("STRIPE_API_KEY")

	// initiate user input reader
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the email address:")

	//call the reader to read the address
	key, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	params := &stripe.CustomerListParams{
		//pass the trimmed whitespaces of the email
		Email: stripe.String(strings.TrimSpace(key)),
	}
	//prevents auto pagination
	params.Single = true
	params.Limit = stripe.Int64(10)

	i := customer.List(params)

	//check to see if customer exists
	if i.Next() {
		c := i.Customer()
		fmt.Println("✅ Customer found: ", c.Name, "-", c.Email)
	} else {
		fmt.Println("❌ No customer found.")
	}

}

// customerCmd represents the config command
var customersCmd = &cobra.Command{
	Use:   "customers",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		getCustomers()
	},
}

func init() {
	rootCmd.AddCommand(customersCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// customersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// customersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
