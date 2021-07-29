package command

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/RafaelYon/kkkj-bot/domain"
	"github.com/RafaelYon/kkkj-bot/repository"
)

type MessagePointsDTO struct {
	Name   string
	Amount int
}

func NewMessagePointsDTO(message string) *MessagePointsDTO {
	temp := strings.Split(message, " ")

	amount, err := strconv.Atoi(temp[2])

	if err != nil {
		fmt.Println("Error in convert amount value")
		fmt.Println(err)
	}

	return &MessagePointsDTO{
		Name:   temp[1],
		Amount: amount,
	}
}

func ContabilizePoints(message string, devRepository repository.DevSaveRetriver) {
	messageDTO := NewMessagePointsDTO(message)

	dev, err := devRepository.Get(messageDTO.Name)

	if err != nil {
		dev = domain.NewDev(messageDTO.Name)
	}

	dev.Add(messageDTO.Amount)

	devRepository.Save(dev)
}
