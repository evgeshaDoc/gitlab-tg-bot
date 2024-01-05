package main

import "github.com/evgeshaDoc/gitlab-tg-bot/internal"

func main() {
	_, err := internal.GetConfig()
	if err != nil {
		panic(err)
	}
}
