package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	telebot "gopkg.in/telebot.v3"
)

var (
	TeleToken = os.Getenv("TELE_TOKEN")
)

var bleblabotCmd = &cobra.Command{
	Use:     "bleblabot",
	Aliases: []string{"start"},
	Short:   "",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf("bleblabot %s started", appVersion)

		bleblabot, err := telebot.NewBot(telebot.Settings{
			URL:    "",
			Token:  TeleToken,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		})

		if err != nil {
			log.Fatalf("Plaese check TELE_TOKEN env variable. %s", err)
			return
		}

		bleblabot.Handle(telebot.OnText, func(m telebot.Context) error {
			log.Print(m.Message().Payload, m.Text())
			payload := m.Message().Payload

			switch payload {
			case "hello":
				err = m.Send(fmt.Sprint("Hello I'm bleblabot!", appVersion))
			}
			return nil
		})

		bleblabot.Start()
	},
}

func init() {
	rootCmd.AddCommand(bleblabotCmd)
}
