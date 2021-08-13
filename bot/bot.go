package bot

import (
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/RomanenkovDV/DechoBot/telegram"
)

type Bot struct {
	Api telegram.Api
}

func (bot *Bot) Serve(resp http.ResponseWriter, req *http.Request) {
	body, _ := io.ReadAll(req.Body)
	update := telegram.ParseUpdate(body)

	if update.Message.Reply.Id != 0 {
		delay, err := time.ParseDuration(update.Message.Text)
		if delay > 0 && err == nil {
			go bot.answer(update.Message.Chat.Id, "You will be notified after "+delay.String())
			go bot.notify(update.Message.Chat.Id, update.Message.Reply.Id, "Message notification", delay)
		} else {
			go bot.answer(update.Message.Chat.Id, "Can't understand the delay, try something like 10s.")
		}
	}
}

func (bot *Bot) answer(chat_id int, text string) {
	id := strconv.Itoa(chat_id)
	bot.Api.Answer(id, text)
}

func (bot *Bot) notify(chat_id int, reply_id int, text string, delay time.Duration) {
	chat := strconv.Itoa(chat_id)
	reply := strconv.Itoa(reply_id)
	<-time.NewTimer(delay).C
	bot.Api.Notify(chat, text, reply)
}
