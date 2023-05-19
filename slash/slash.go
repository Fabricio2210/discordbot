package slash

import (
	"discordmarlin/random"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"strings"
	"os"
)

// MessageCommands handles message commands.
// It checks if the message contains any words and performs actions accordingly.
func MessageCommands(s *discordgo.Session, m *discordgo.MessageCreate, bussWords []string, command string) error {
	// Load environment variables from "local.env" file
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatalf("Some error occurred. Err: %s", err)
	}

	// Check each word in the message content
	for _, keyword := range bussWords {
		if strings.Contains(strings.ToLower(m.Content), keyword) {
			// Delete the triggering message
			err = s.ChannelMessageDelete(m.ChannelID, m.ID)
			if err != nil {
				fmt.Println("Error deleting message:", err)
			}

			// Get the links associated with the command from environment variables
			links := strings.Split(os.Getenv(command), ",")

			// Choose a random link from the available links
			randomLink := random.ChooseRandomItem(links)

			// Construct the message containing the random link
			msg := randomLink

			// Send the message with the random link to the channel
			_, err := s.ChannelMessageSend(m.ChannelID, msg)
			if err != nil {
				fmt.Println("Error sending message:", err)
			}
		}
	}
	return nil
}