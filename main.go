package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/qepo17/go-openrouter"
	"github.com/urfave/cli/v2"
)

func main() {
	cfg, err := GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	openRouterDomain, err := openrouter.New(cfg.OpenRouterAPIKey, openrouter.ClientOptions{})
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	app := App(ctx, openRouterDomain)

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func App(ctx context.Context, openRouterDomain *openrouter.Client) *cli.App {
	return &cli.App{
		Name:  "OpenAI CLI",
		Usage: "Ask questions to OpenAI's GPT-3",
		Commands: []*cli.Command{
			{
				Name:    "ask",
				Aliases: []string{"a"},
				Usage:   "Ask a question to OpenAI's GPT-3",
				Action: func(c *cli.Context) error {
					question := c.Args().First()

					resp, err := openRouterDomain.Completions(ctx, openrouter.CompletionsRequest{
						Prompt: question,
					})
					if err != nil {
						return err
					}

					answer := ""
					for _, choice := range resp.Choices {
						answer += choice.Text + "\n"
					}

					fmt.Println(answer)

					return nil
				},
			},
		},
	}
}
