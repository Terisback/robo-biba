package main

import (
	"context"
	"log"

	"github.com/andersfylling/disgord"
	"github.com/codingconcepts/env"
	"github.com/jonas747/dca"
	"github.com/rylio/ytdl"
)

const (
	RickRollURL = "https://www.youtube.com/watch?v=52Gg9CqhbP8&ab_channel=STUCKINTHESOUND"
)

type Config struct {
	Token string `env:"TOKEN" required:"true"`
}

var (
	config Config
)

func init() {
	if err := env.Set(&config); err != nil {
		log.Fatalln("Error caused by trying to get environment variables,", err)
	}
}

func main() {
	dg := disgord.New(disgord.Config{
		BotToken: config.Token,
	})
	defer dg.StayConnectedUntilInterrupted(context.Background())

	var voice disgord.VoiceConnection
	dg.Ready(func() {
		// Once the bot has connected to the websocket, also connect to the voice channel
		voice, _ = dg.VoiceConnect(disgord.NewSnowflake(701878021421793330), disgord.NewSnowflake(701878021988155396))
	})
	dg.On(disgord.EvtMessageCreate, func(_ disgord.Session, m *disgord.MessageCreate) {
		if m.Message.Content == "!what" {
			options := dca.StdEncodeOptions
			options.RawOutput = true
			options.Bitrate = 96
			options.Application = "audio"

			videoInfo, err := ytdl.GetVideoInfo(context.Background(), RickRollURL)
			if err != nil {
				log.Fatalln(err)
			}

			format := videoInfo.Formats.Extremes(ytdl.FormatAudioBitrateKey, true)[0]
			downloadURL, err := ytdl.DefaultClient.GetDownloadURL(context.Background(), videoInfo, format)
			if err != nil {
				log.Fatalln(err)
			}

			encodingSession, err := dca.EncodeFile(downloadURL.String(), options)
			if err != nil {
				log.Fatalln(err)
			}
			defer encodingSession.Cleanup()

			_ = voice.StartSpeaking()
			_ = voice.SendDCA(encodingSession)
			_ = voice.StopSpeaking()
		}
	})
}
