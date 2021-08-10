package telegram

type Update struct {
	Id      int     `json:"update_id"`
	Message Message `json:"message"`
}

type Message struct {
	Id    int    `json:"message_id"`
	Text  string `json:"text"`
	Chat  Chat   `json:"chat"`
	Reply Reply  `json:"reply_to_message"`
}

type Chat struct {
	Id int `json:"id"`
}

type Reply struct {
	Id int `json:"message_id"`
}
