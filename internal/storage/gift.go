package storage

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"time"
)

const GiftPrefix = "gift"

func GiftStatus(guildID, userID string) (gained bool, expiration time.Time, err error) {
	key := []byte(GiftPrefix + guildID + userID)

	exist, err := db.Has(key)
	if err != nil {
		return false, time.Time{}, err
	}

	if exist {
		data, err := db.Get(key)
		if err != nil {
			return false, time.Time{}, err
		}

		var (
			rw bytes.Buffer
			gb time.Time
		)
		dec := gob.NewDecoder(&rw)

		_, err = rw.Write(data)
		if err != nil {
			return false, time.Time{}, err
		}

		err = dec.Decode(&gb)
		if err != nil {
			return false, time.Time{}, err
		}

		if time.Now().UTC().Before(gb) {
			return true, gb, nil
		} else {
			return false, time.Time{}, nil
		}
	}

	return false, time.Time{}, nil
}

func AddGiftBound(guildID, userID string, expiration time.Time) (err error) {
	key := []byte(GiftPrefix + guildID + userID)

	var rw bytes.Buffer
	enc := gob.NewEncoder(&rw)

	err = enc.Encode(expiration)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(&rw)
	if err != nil {
		return err
	}

	err = db.Put(key, data)
	if err != nil {
		return err
	}

	return nil
}
