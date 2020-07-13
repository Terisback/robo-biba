package storage

import (
	"log"
	"time"

	"github.com/akrylysov/pogreb"
)

var (
	db *pogreb.DB
)

func Init() {
	var err error
	db, err = pogreb.Open("db", &pogreb.Options{
		BackgroundCompactionInterval: time.Minute * 5,
		BackgroundSyncInterval:       -1, // Doing sync every write operation
	})
	if err != nil {
		log.Fatalln(err)
	}
}

func Close() {
	if err := db.Close(); err != nil {
		log.Fatalln(err)
	}
}
