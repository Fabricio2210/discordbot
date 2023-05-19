package chatlogs

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"time"
	"path/filepath"

)

func Chatlogs(s *discordgo.Session, m *discordgo.MessageCreate) {
	layout := "2006-01-02"
	currentTime := time.Now()
	dateString := currentTime.Format(layout)
	fileDayPath := filepath.Join("C:\\Users\\fabri\\OneDrive\\Área de Trabalho\\drive\\dtraktorAgregrator\\superchatExc", dateString+"_superchats_day.xlsx")
	fileDay, err := os.Open(fileDayPath)
	if err != nil {
		fmt.Println("File not found:", fileDayPath)
		return
	}
	defer fileDay.Close()
	fileNightPath := filepath.Join("C:\\Users\\fabri\\OneDrive\\Área de Trabalho\\drive\\dtraktorAgregrator\\superchatExc", dateString+"_superchats_night.xlsx")
	fileNight, err := os.Open(fileNightPath)
	if err != nil {
		fmt.Println("File not found:", fileNightPath)
		return
	}
	defer fileNight.Close()

	dFileDay := &discordgo.File{
		Name:        fmt.Sprintf("./%s_superchats_day.xlsx", dateString),
		Reader:      fileDay,
		ContentType: "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
	}
	dFileNight := &discordgo.File{
		Name:        fmt.Sprintf("./%s_superchats_night.xlsx", dateString),
		Reader:      fileNight,
		ContentType: "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
	}
	files := []*discordgo.File{dFileDay, dFileNight}
	_, err = s.ChannelMessageSendComplex("1105617127915933796", &discordgo.MessageSend{
		Files: files,
	})
	if err != nil {
		log.Fatal("Error sending message:", err)
		return
	}
	fmt.Println("Files sent successfully!")
}
