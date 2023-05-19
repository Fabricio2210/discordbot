package main

import (
	//"discordmarlin/automessages"
	"discordmarlin/cron"
	"discordmarlin/message"
	"discordmarlin/slash"
	"discordmarlin/whoasked"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	// Create a new Discord session
	discord, err := discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		fmt.Println("Error creating Discord session:", err)
		return
	}

	if err != nil {
		panic(err)
	}

	// err = automessages.ScheduleMessage(discord)
	// if err != nil {
	// 	fmt.Println("Error scheduling message:", err)
	// 	return
	// }
	go cron.RunChatlogsAtScheduledTime(discord)
	discord.AddHandler(whoasked.MessageWhoAsked)
	discord.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		message.MessageCreate(s, m, os.Getenv("MARLIN") , []string{"Brazilian"}, "STFU MARLIN!!ʚ♡ɞʚ♡ɞʚ♡ɞʚ♡ɞ")
	})
	discord.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		slash.MessageCommands(s, m, []string{"/buss", "/ambasing"}, "BUSS")
	})
	discord.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		slash.MessageCommands(s, m, []string{"/derich", "/derish", "/pervert"}, "DERICH")
	})
	discord.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		slash.MessageCommands(s, m, []string{"/jowls", "/moobs"}, "JOWLS")
	})
	err = discord.Open()
	if err != nil {
		fmt.Println("Error opening connection:", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	<-make(chan struct{})
}
