package command

import "github.com/bwmarrin/discordgo"

func Help(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "```"+`
--------------------------------------------
π₯ ;points [name] [amount]
Add pontos para alguem
@name string
@amount int
--------------------------------------------
π ;polls
Listagem da votaΓ§Γ£o atual
--------------------------------------------
ππ½ ;help
Ajuda nΓ© 
--------------------------------------------
π ;pong
P I N G
--------------------------------------------
π³οΈβπ ;oh dad
STOP
--------------------------------------------
π οΈ ;staging
Adiciona mais um na quantidade de vezes que recriamos a staging
--------------------------------------------
`+"```")
}
