package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "new",
		Usage: "make a new command with the Cli package",
		Action: func(*cli.Context) error {
			fmt.Println("Here's a new command")
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
