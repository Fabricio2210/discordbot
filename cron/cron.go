package cron

import (
	"time"

	"github.com/bwmarrin/discordgo"
	"discordmarlin/chatlogs"
)

func RunChatlogsAtScheduledTime(s *discordgo.Session) {
	// Get the current time
	now := time.Now()
	// Calculate the duration until the next 7 PM
	next7PM := time.Date(now.Year(), now.Month(), now.Day(), 19, 0, 0, 0, now.Location())
	if now.After(next7PM) {
		next7PM = next7PM.Add(24 * time.Hour) // Move to the next day
	}
	durationUntilNext7PM := next7PM.Sub(now)

	// Wait until the next 7 PM
	time.Sleep(durationUntilNext7PM)

	// Run Chatlogs function
	chatlogs.Chatlogs(s, nil,"C:\\Users\\Third\\Desktop\\dtraktorAgregrator\\superchatExc", "1105617127915933796")

	// Schedule Chatlogs function to run every day at 7 PM
	ticker := time.NewTicker(24 * time.Hour)
	for range ticker.C {
		chatlogs.Chatlogs(s, nil, "C:\\Users\\Third\\Desktop\\dtraktorAgregrator\\superchatExc", "1105617127915933796")
	}
}