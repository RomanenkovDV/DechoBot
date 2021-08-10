package telegram

import "encoding/json"

///
type Api interface {
	TokenCheck() bool
	SetWebhook(url string) bool
	Answer(chat_id string, text string) bool
	Notify(chat_id string, text string, reply_id string) bool
}

//
func ParseUpdate(data []byte) Update {
	var update Update
	json.Unmarshal(data, &update)
	return update
}
