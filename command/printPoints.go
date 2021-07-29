package command

import (
	"strconv"

	"github.com/RafaelYon/kkkj-bot/repository"
	"github.com/bwmarrin/discordgo"
)

func PrintPoints(s *discordgo.Session, m *discordgo.MessageCreate, devRepository *repository.Dev) {
	template := "```[\n"

	for _, v := range devRepository.GetAll() {
		template += v.Name + " => " + strconv.Itoa(v.Amount) + "\n"
	}

	template += "]```"

	s.ChannelMessageSend(m.ChannelID, template)
}
