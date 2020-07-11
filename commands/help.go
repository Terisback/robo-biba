package commands

import (
	"context"
	"fmt"

	"github.com/andersfylling/disgord"
)

func Help(session disgord.Session, event *disgord.MessageCreate) {
	var help string
	help += "`!онлайн` - Статистика по онлайну\n"
	help += "`!онлайн <id или упоминание роли>` - Статистика онлайна роли\n"
	help += "`!когда <я, id или упоминание юзера>` - Когда зашёл на сервер\n"

	embed := disgord.Embed{
		Title:       "Список команд",
		Description: help,
	}

	_, err := session.SendMsg(context.Background(), event.Message.ChannelID, embed)
	if err != nil {
		fmt.Println(err)
	}
}
