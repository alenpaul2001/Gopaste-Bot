package main

import (
	"fmt"
	"log"
	"os"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	keyboard  = &tb.ReplyMarkup{}
	nekoboard = keyboard.Data("nekobin", "neko")
	dogboard  = keyboard.Data("dogbin", "dog")
)

func main() {
	b, err := tb.NewBot(tb.Settings{
		Token:     os.Getenv("BOT_TOKEN"),
		Poller:    &tb.LongPoller{Timeout: 10 * time.Second},
		ParseMode: "markdown",
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	keyboard.Inline(keyboard.Row(nekoboard, dogboard))

	b.Handle("/gopaste", func(m *tb.Message) {
		if m.Chat.Type == "group" || m.Chat.Type == "supergroup" {
			b.Send(
				m.Chat, "**please choose a paste service**", &tb.SendOptions{
					ReplyTo:     m,
					ReplyMarkup: keyboard,
				},
			)
		}

	})

	b.Handle(&nekoboard, func(c *tb.Callback) {
		b.Edit(c.Message, "clicked nekobin")
	})

	b.Handle(&nekoboard, func(c *tb.Callback) {
		b.Edit(c.Message, "clicked dogbin")
	})

	fmt.Println("Bot is now running")
	b.Start()
}
