package main

import (
	"github.com/polRk/telegram"
	"log"
	"os"
)

func main() {
	token := os.Getenv("TOKEN")

	client := telegram.NewClient(token)
	bot, err := client.GetMe()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(bot)
}
