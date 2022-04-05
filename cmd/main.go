package main

import (
	"github.com/polRk/telegram/bot/api"
	"log"
	"os"
)

func main() {
	token := os.Getenv("TOKEN")

	client := api.NewClient(token)
	bot, err := client.GetMe()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(bot)
}
