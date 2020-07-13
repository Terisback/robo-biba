package storage

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
)

const CurrencyPrefix = "currency"

func Balance(guildID, userID string) (balance int, err error) {
	key := []byte(CurrencyPrefix + guildID + userID)

	exist, err := db.Has(key)
	if err != nil {
		return 0, err
	}

	if exist {
		data, err := db.Get(key)
		if err != nil {
			return 0, err
		}

		var rw bytes.Buffer
		dec := gob.NewDecoder(&rw)

		_, err = rw.Write(data)
		if err != nil {
			return 0, err
		}

		err = dec.Decode(&balance)
		if err != nil {
			return 0, err
		}

		return balance, nil
	}

	return 0, nil
}

func AddCurrency(guildID, userID string, amount int) (balance int, err error) {
	key := []byte(CurrencyPrefix + guildID + userID)

	exist, err := db.Has(key)
	if err != nil {
		return 0, err
	}

	if exist {
		balance, err = Balance(guildID, userID)
		if err != nil {
			return 0, err
		}

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

		return balance + amount, nil
	}

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

func SubCurrency(guildID, userID string, amount int) (balance int, err error) {
	return AddCurrency(guildID, userID, -amount)
}
