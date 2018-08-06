package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "trader"
	app.Usage = "lets get rich"

	app.Commands = []cli.Command{
		{
			Name:    "earnings",
			Aliases: []string{"e"},
			Usage:   "check stocks which have had earnings releases recently",
			Action:  fetchEarnings,
		},
		{
			Name:    "quote",
			Aliases: []string{"q"},
			Usage:   "get a stock quote",
			Action:  fetchQuotes,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func fetchQuotes(c *cli.Context) error {
	fmt.Println("fetching quote")
	for _, q := range c.Args() {
		fmt.Println(q)
	}
	return nil
}

func fetchEarnings(c *cli.Context) error {
	fmt.Println("checking earnings")
	return nil
}
