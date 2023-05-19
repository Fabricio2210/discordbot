package message

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

// MessageCreate handles the creation of new messages.
// It checks if the message author's ID matches the provided ID and if the message contains any blacklisted keywords.
// If a blacklisted keyword is found, it deletes the message and sends a specified response message.
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate, Id string, blacklist []string, msg string) error {
	fmt.Println("channel id:", m.ChannelID)
	fmt.Println("content:", m.Content)

	// Check if the author's ID matches the provided ID
	if m.Author.ID == Id {
		// Check each blacklisted keyword in the message content
		for _, keyword := range blacklist {
			if strings.Contains(strings.ToLower(m.Content), keyword) {
				// Delete the triggering message
				err := s.ChannelMessageDelete(m.ChannelID, m.ID)
				if err != nil {
					fmt.Println("Error deleting message:", err)
				}

				// Send the specified response message to the channel
				_, err = s.ChannelMessageSend(m.ChannelID, msg)
				if err != nil {
					fmt.Println("Error sending message:", err)
				}
			}
		}
	}
	return nil
}
