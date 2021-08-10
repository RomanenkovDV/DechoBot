package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/RomanenkovDV/DechoBot/bot"
	"github.com/RomanenkovDV/DechoBot/ngrok"
	"github.com/RomanenkovDV/DechoBot/telegram"
)

func main() {
	log.Println("Attempt to start...")

	// handle command line args
	var (
		flagHost        = flag.String("host", "127.0.0.1:8080", "127.0.0.1:8080")
		flagToken       = flag.String("token", "", "telegram_bot_token")
		flagWebhookUrl  = flag.String("hook", "https://127.0.0.1:8080", "ngrok, https://127.0.0.1:8080")
		flagNgrokWebIfc = flag.String("nweb", "http://127.0.0.1:4040", "ngrok web interface url")
	)
	flag.Parse()

	var (
		host              = *flagHost
		token             = *flagToken
		webhookUrl        = *flagWebhookUrl
		ngrokWebInterface = *flagNgrokWebIfc
	)

	// get tunnel url if ngrok is using
	if webhookUrl == "ngrok" {
		tunnel, success := ngrok.FetchTunnelAddress(ngrokWebInterface)
		if !success {
			log.Fatalln("Unable to fetch webhook url.")
		}
		log.Println("Using ngrok tunnel on:", tunnel)
		webhookUrl = tunnel
	}

	// check token
	api := &telegram.ApiImplNative{Token: token}
	if !api.TokenCheck() {
		log.Fatal("Token is invalid")
	}
	log.Println("Token is ok")

	// set webhook
	if !api.SetWebhook(webhookUrl) {
		log.Fatal("Can not to set webhook on:", webhookUrl)
	}
	log.Println("Webhook was set successfully")

	/// start server
	bot := bot.Bot{Api: api}
	http.HandleFunc("/", bot.Serve)
	log.Println("Server starting on:", host)
	log.Fatal(http.ListenAndServe(host, nil))
}
