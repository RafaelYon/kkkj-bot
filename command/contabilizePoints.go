package command

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/RafaelYon/kkkj-bot/domain"
	"github.com/RafaelYon/kkkj-bot/repository"
	"github.com/bwmarrin/discordgo"
)

type MessagePointsDTO struct {
	Name   string
	Amount int
}

func NewMessagePointsDTO(message string) (*MessagePointsDTO, error) {
	temp := strings.Split(message, " ")

	if len(temp) < 3 {
		return nil, fmt.Errorf("BURRO BURRO BURRO")
	}

	amount, err := strconv.Atoi(temp[2])

	if err != nil {
		fmt.Println("Error in convert amount value")
		fmt.Println(err)
	}

	return &MessagePointsDTO{
		Name:   temp[1],
		Amount: amount,
	}, nil
}

func ContabilizePoints(s *discordgo.Session, m *discordgo.MessageCreate, devRepository repository.DevSaveRetriver) error {
	messageDTO, err := NewMessagePointsDTO(m.Content)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, err.Error())
		return err
	}

	dev, err := devRepository.Get(messageDTO.Name)

	if err != nil {
		dev = domain.NewDev(messageDTO.Name)
	}

	dev.Add(messageDTO.Amount)

	devRepository.Save(dev)
	return nil
}
