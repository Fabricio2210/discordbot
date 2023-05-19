package whoasked

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func MessageWhoAsked(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == "302205824011337729" {
		blacklist := []string{"faith", "jew","god", "jesus","christ", "reality", "miracles", "church", "niggers", "brazilian", "testament"}
		for _, keyword := range blacklist {
			if strings.Contains(strings.ToLower(m.Content), keyword) {
				msg := "https://media.tenor.com/2Gsf2UQ7Qw4AAAAC/who-asked.gif"
				_, err := s.ChannelMessageSend(m.ChannelID, msg)
				if err != nil {
					fmt.Println("error sending message,", err)
				}
				return
			}
		}
	}
}
