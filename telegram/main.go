package main

import "log"
import "github.com/go-telegram-bot-api/telegram-bot-api"

func main() {
	bot, err := tgbotapi.NewBotAPI("447466867:AAGE4dx4UzmdbwCuh5lmCuPVl245COQ73Tk")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 100000

	updates, err := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message.Text == "/keyup" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Select One")
			var s *string
			*s = "ok"
			msg.ReplyMarkup = tgbotapi.InlineKeyboardMarkup{
				InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
					[]tgbotapi.InlineKeyboardButton{
						tgbotapi.InlineKeyboardButton{Text: "aa", CallbackData: s},
						tgbotapi.InlineKeyboardButton{Text: "bb", CallbackData: s},
					},
				},
			}
			reMsg, err := bot.Send(msg)
			if err != nil {
				log.Fatalln(err)
			}
			log.Printf("[%s] %s", reMsg.From.UserName, reMsg.Text)
		}
	}
}
