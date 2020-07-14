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
	help += "`!баланс` - Узнать свой баланс\n"
	help += "`!подарок` - Получить в подарок раз в 2 часа от 10 до 50 монет\n"
	help += "`!флип <сумма>` - Подкинуть монетку\n"

	embed := disgord.Embed{
		Title:       "Список команд",
		Description: help,
		Color:       getIntColor(defaultEmbedColor),
	}

	_, err := event.Message.Reply(context.Background(), session, &embed)
	if err != nil {
		fmt.Println(err)
	}
}
