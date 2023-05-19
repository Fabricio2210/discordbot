package automessages

import (
	"discordmarlin/random"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"math/rand"
	"time"
)

func ScheduleMessage(dg *discordgo.Session) error {
	go func() {
		for {
			rand.Seed(time.Now().UnixNano())
			duration := time.Duration(rand.Intn(11000)+1100) * time.Second // generate random duration between 500 and 6000 seconds
			fmt.Println("the number is", duration)
			ticker := time.NewTicker(duration)
			<-ticker.C
			messages := []string{"Hope you're having a wonderful day", "I love you Colton!", "I respect the Jewish people and their struggles", "I love Amber!", "Jamie? Who is Jamie?", "Black lives matter"}
			randomMessage := random.ChooseRandomItem(messages)
			msg := randomMessage
			_, err := dg.ChannelMessageSend("909966538151448589", msg)
			if err != nil {
				fmt.Println("Error sending message:", err)
				return
			}
		}

	}()
	return nil
}
//messages := []string{"My dick doesnt work!", "I miss Jamie", "I love my life!", "Sucks to suck", "Shut up jew!", "Shut up brazilian", "Fuck you Colton!", "Nigger mike fucked me in the ass last night", "Goddamn Marlin, you ain't worth a fuck!"} //"https://i.imgur.com/P6mkjoE.jpg"