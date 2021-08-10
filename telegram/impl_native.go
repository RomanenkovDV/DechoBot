package telegram

import (
	"log"
	"net/http"
)

const (
	c_HOST             = "https://api.telegram.org/bot"
	c_MTH_GET_ME       = "getMe"
	c_MTH_SET_WEBHOOK  = "setWebhook"
	c_MTH_SEND_MESSAGE = "sendMessage"
)

type ApiImplNative struct {
	Token string
}

//
func (api *ApiImplNative) call(method string) string {
	return c_HOST + api.Token + "/" + method + "?"
}

//
func (api *ApiImplNative) TokenCheck() bool {
	status := false
	resp, err := http.Get(api.call(c_MTH_GET_ME))

	if resp != nil {
		defer resp.Body.Close()
		status = resp.StatusCode == http.StatusOK
	} else {
		if err != nil {
			log.Println(err)
		}
	}

	return status
}

//
func (api *ApiImplNative) SetWebhook(url string) bool {
	status := false
	resp, err := http.Get(api.call(c_MTH_SET_WEBHOOK) + "url=" + url)

	if resp != nil {
		defer resp.Body.Close()
		status = resp.StatusCode == http.StatusOK
	} else {
		if err != nil {
			log.Println(err)
		}
	}

	return status
}

//
func (api *ApiImplNative) Answer(chat_id string, text string) bool {
	request := api.call(c_MTH_SEND_MESSAGE) +
		"chat_id=" + chat_id + "&" +
		"text=" + text

	status := false
	resp, err := http.Get(request)
	if resp != nil {
		defer resp.Body.Close()
		status = resp.StatusCode == http.StatusOK
	} else {
		if err != nil {
			log.Println(err)
		}
	}

	return status
}

//
func (api *ApiImplNative) Notify(chat_id string, text string, reply_id string) bool {
	request := api.call(c_MTH_SEND_MESSAGE) +
		"chat_id=" + chat_id + "&" +
		"text=" + text + "&" +
		"reply_to_message_id=" + reply_id + "&" +
		"allow_sending_without_reply=True"

	status := false
	resp, err := http.Get(request)
	if resp != nil {
		defer resp.Body.Close()
		status = resp.StatusCode == http.StatusOK
	} else {
		if err != nil {
			log.Println(err)
		}
	}

	return status
}
