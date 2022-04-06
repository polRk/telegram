package main

import (
	"context"
	"github.com/polRk/telegram"
	"log"
	"os"
)

func main() {
	token := os.Getenv("TOKEN")
	client := telegram.New(token)

	bot, err := client.GetMe(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(bot)
}
