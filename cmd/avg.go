package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var avgCmd = &cobra.Command{
	Use:   "avg",
	Short: "A brief description of the deepsubcommand",
	Long: `A longer description that explains the deepsubcommand
in detail. For example, you can mention how it complements
the main command and what options it supports.`,
	Run: func(cmd *cobra.Command, args []string) {
		// This function will be executed when the "deepsubcommand" is called
		fmt.Println("MRR averages: ")
	},
}

func init() {
	//avgCmd is a deepsubcommand of mrrCmd
	mrrCmd.AddCommand(avgCmd)
}
