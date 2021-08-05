package command

import "github.com/bwmarrin/discordgo"

func Help(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "```"+`
--------------------------------------------
🔥 ;points [name] [amount]
Add pontos para alguem
@name string
@amount int
--------------------------------------------
🏆 ;polls
Listagem da votação atual
--------------------------------------------
🙏🏽 ;help
Ajuda né 
--------------------------------------------
🏓 ;pong
P I N G
--------------------------------------------
🏳️‍🌈 ;oh dad
STOP
--------------------------------------------
🛠️ ;staging
Adiciona mais um na quantidade de vezes que recriamos a staging
--------------------------------------------
`+"```")
}
