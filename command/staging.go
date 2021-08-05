package command

import (
	"strconv"

	"github.com/RafaelYon/kkkj-bot/repository"
	"github.com/bwmarrin/discordgo"
)

func Staging(r *repository.Staging, s *discordgo.Session, m *discordgo.MessageCreate) {
	r.AddAmount()
	s.ChannelMessageSend(m.ChannelID, "Está é a "+strconv.Itoa(r.Amount)+"º que recriamos a staging")
}
