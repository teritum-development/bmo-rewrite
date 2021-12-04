package main

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

var discord *discordgo.Session
var botID string

func init() {
	var err error

	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		panic("DISCORD_TOKEN unset")
	}
	log.Println("Logging in...")
	discord, err = discordgo.New("Bot " + token)
	if err != nil {
		panic(err)
	}

	discord.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMembers | discordgo.IntentsAllWithoutPrivileged)

	user, err := discord.User("@me")
	if err != nil {
		panic(err)
	}

	botID = user.ID
	log.Println("Logging in with", botID)

}

func main() {
	err := discord.Open()
	if err != nil {
		panic(err)
	}
	log.Println("Connected to discord")
	forever := make(chan int)
	<-forever
}
