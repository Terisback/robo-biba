package economy

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/akrylysov/pogreb"
)

var (
	db *pogreb.DB
)

func init() {
	var err error
	db, err = pogreb.Open("economy.db", &pogreb.Options{BackgroundSyncInterval: time.Minute * 1, BackgroundCompactionInterval: time.Minute * 10})
	if err != nil {
		log.Fatal(err)
		return
	}
}

func Close() {
	err := db.Sync()
	if err != nil {
		fmt.Println(err)
	}
	db.Close()
}

func Balance(guildID, userID string) (balance int, err error) {
	key := []byte(guildID + userID)
	alreadyInDB, err := db.Has(key)
	if err != nil {
		return 0, err
	}

	if !alreadyInDB {
		return 0, nil
	}

	var rw bytes.Buffer
	dec := gob.NewDecoder(&rw)
	data, err := db.Get(key)
	if err != nil {
		return 0, err
	}
	rw.Write(data)

	err = dec.Decode(&balance)
	if err != nil {
		return 0, err
	}

	return balance, nil
}

func Add(guildID, userID string, amount int) (balance int, err error) {
	key := []byte(guildID + userID)

	alreadyInDB, err := db.Has(key)
	if err != nil {
		return 0, err
	}

	switch alreadyInDB {
	case true:
		var rw bytes.Buffer
		enc := gob.NewEncoder(&rw)
		dec := gob.NewDecoder(&rw)

		data, err := db.Get(key)
		if err != nil {
			return 0, err
		}
		rw.Write(data)

		err = dec.Decode(&balance)
		if err != nil {
			return 0, err
		}
		balance += amount

		err = enc.Encode(balance)
		if err != nil {
			return 0, err
		}

		data, err = ioutil.ReadAll(&rw)
		if err != nil {
			return 0, err
		}

		err = db.Put(key, data)
		if err != nil {
			return 0, err
		}

		return balance, nil
	case false:
		var rw bytes.Buffer
		enc := gob.NewEncoder(&rw)

		err = enc.Encode(amount)
		if err != nil {
			return 0, err
		}

		data, err := ioutil.ReadAll(&rw)
		if err != nil {
			return 0, err
		}

		err = db.Put(key, data)
		if err != nil {
			return 0, err
		}

		return amount, nil
	}

	panic("Unreachable!")
}
