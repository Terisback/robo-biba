package commands

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/Terisback/robo-biba/economy"
	"github.com/andersfylling/disgord"
	"github.com/patrickmn/go-cache"
)

const (
	giftCacheFilename        = "giftCache.dat"
	giftAlreadyGiftedMessage = "%s вы уже получили подарок!"
	giftNewGift              = "%s поздравляю вы получили %d монет!"
	giftAlreadyDesc          = "До следующего подарка осталось %s\n__Баланс:__ **%d** <:rgd_coin_rgd:518875768814829568>"
	giftDesc                 = "Следующий подарок можно будет получить через 2 часа!\n__Баланс:__ **%d** <:rgd_coin_rgd:518875768814829568>"
)

var (
	giftCache = cache.New(time.Hour*2, time.Minute*5)
	giftRand  = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func loadCache() {
	data, err := ioutil.ReadFile(giftCacheFilename)
	if err != nil {
		log.Fatalln(err)
		return
	}

	var rw bytes.Buffer
	dec := gob.NewDecoder(&rw)

	_, err = rw.Write(data)
	if err != nil {
		log.Fatalln(err)
		return
	}

	items := make(map[string]cache.Item)
	err = dec.Decode(&items)
	if err != nil {
		log.Fatalln(err)
		return
	}

	giftCache = cache.NewFrom(time.Hour*2, time.Minute*1, items)
}

func saveCache() {
	items := giftCache.Items()

	var rw bytes.Buffer
	enc := gob.NewEncoder(&rw)

	err := enc.Encode(items)
	if err != nil {
		log.Fatalln(err)
		return
	}

	data, err := ioutil.ReadAll(&rw)
	if err != nil {
		log.Fatalln(err)
		return
	}

	err = ioutil.WriteFile(giftCacheFilename, data, 0664)
	if err != nil {
		log.Fatalln(err)
		return
	}
}

func init() {
	gob.Register(cache.Item{})

	if _, err := os.Stat("./" + giftCacheFilename); err == nil {
		loadCache()
	} else {
		saveCache()
	}
}

func GiftCacheSave() {
	saveCache()
}

func Gift(session disgord.Session, event *disgord.MessageCreate) {
	guildID := event.Message.GuildID.String()
	userID := event.Message.Author.ID.String()

	var nickname string = event.Message.Member.Nick

	if nickname == "" {
		nickname = event.Message.Author.Username
	}

	avatarURL, err := event.Message.Author.AvatarURL(64, false)
	if err != nil {
		return
	}

	embed := disgord.Embed{}

	if _, expiration, exists := giftCache.GetWithExpiration(guildID + userID); exists {
		d := expiration.UTC().Sub(time.Now().UTC())

		since := durString(d)

		balance, err := economy.Balance(event.Message.GuildID.String(), event.Message.Author.ID.String())
		if err != nil {
			return
		}

		embed.Author = &disgord.EmbedAuthor{IconURL: avatarURL, Name: fmt.Sprintf(giftAlreadyGiftedMessage, nickname)}
		embed.Description = fmt.Sprintf(giftAlreadyDesc, since, balance)
	} else {
		sum := 10 + giftRand.Intn(40)
		balance, err := economy.Add(event.Message.GuildID.String(), event.Message.Author.ID.String(), sum)
		if err != nil {
			return
		}

		giftCache.SetDefault(guildID+userID, true)

		embed.Author = &disgord.EmbedAuthor{IconURL: avatarURL, Name: fmt.Sprintf(giftNewGift, nickname, sum)}
		embed.Description = fmt.Sprintf(giftDesc, balance)
	}

	session.SendMsg(context.Background(), event.Message.ChannelID, &embed)
}

func durString(d time.Duration) (result string) {
	var seconds, minutes, hours int
	seconds = int(d.Seconds())
	if seconds > 60 {
		minutes = (seconds - seconds%60) / 60
	}
	if minutes > 59 {
		hours = (minutes - minutes%60) / 60
		minutes -= hours * 60
		result = strconv.Itoa(hours)
		result += " " + hoursTail(hours)
	}
	if minutes != 0 {
		if result != "" {
			result += " "
		}
		result += strconv.Itoa(minutes)
		result += " " + minutesTail(minutes)
	}
	return
}

func hoursTail(hours int) (result string) {
	switch {
	case hours > 20 && hours < 100:
		hours %= 10
	case hours > 100:
		hours = hours % 100 % 10
	}
	switch hours {
	case 1:
		result = "час"
	case 2, 3, 4:
		result = "часа"
	default:
		result = "часов"
	}
	return
}

func minutesTail(minutes int) (result string) {
	switch {
	case minutes > 20 && minutes < 60:
		minutes %= 10
	case minutes > 60:
		minutes = minutes % 60 % 10
	}
	switch minutes {
	case 1:
		result = "минута"
	case 2, 3, 4:
		result = "минуты"
	default:
		result = "минут"
	}
	return
}
