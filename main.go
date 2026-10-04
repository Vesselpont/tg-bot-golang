package main

import (
	"flag"
	"golang-lod/clients/telegram"
	"log"
)

const (
	tgBotHost = "ыыыы"
)

func main() {

	tgClient = telegram.New(mustToken())

	//fetcher = fetcher.New

	//proccesor

	//consumer.Start(fetcher, processor)
}

func mustToken() string {
	token := flag.String(
		"token-bot-token",
		"",
		"token for access to telegram")
	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specifid")
	}
	return *token
}
