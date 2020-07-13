package storage

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client     *mongo.Client
	discord    *mongo.Database
	finance    *mongo.Collection
	giftBounds *mongo.Collection
)

func Init(mongoURI string) {
	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	discord = client.Database("discord")
	finance = discord.Collection("finance")
	giftBounds = discord.Collection("gift-bounds")
}

func Close() {
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}

func Balance(guildID, userID string) (balance int, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := guildID + userID
	filter := bson.M{"_id": id}
	bl := finance.FindOne(ctx, filter)

	if bl.Err() != nil {
		switch bl.Err() {
		case mongo.ErrNoDocuments:
			return 0, nil
		default:
			return 0, err
		}
	}

	err = bl.Decode(&balance)
	if err != nil {
		return 0, err
	}

	return balance, nil
}

func AddCurrency(guildID, userID string, amount int) (balance int, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := guildID + userID
	filter := bson.M{"_id": id}
	bl := finance.FindOne(ctx, filter)

	if err = bl.Decode(&balance); err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			ctx, cancelOp := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancelOp()
			_, err = finance.InsertOne(ctx, bson.M{"_id": id, "currency": amount})
			if err != nil {
				return 0, err
			}
			return amount, nil
		default:
			return 0, err
		}
	}

	newBalance := balance + amount

	ctx, cancelBip := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelBip()
	_, err = finance.ReplaceOne(ctx, filter, bson.M{"_id": id, "currency": newBalance})
	if err != nil {
		return 0, err
	}

	return newBalance, nil
}

func SubCurrency(guildID, userID string, amount int) (balance int, err error) {
	return AddCurrency(guildID, userID, -amount)
}

func GiftStatus(guildID, userID string) (gained bool, expiration time.Time, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := guildID + userID
	filter := bson.M{"_id": id}
	gb := giftBounds.FindOne(ctx, filter)

	type res struct {
		Gained     bool
		Expiration []byte
	}

	var r res

	if err = gb.Decode(&r); err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return false, time.Time{}, nil
		default:
			return false, time.Time{}, err
		}
	}

	var tm time.Time

	err = tm.GobDecode(r.Expiration)
	if err != nil {
		return false, time.Time{}, err
	}

	if time.Now().UTC().Before(tm) {
		return r.Gained, tm, nil
	}

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = giftBounds.DeleteOne(ctx, filter)
	if err != nil {
		return false, time.Time{}, err
	}

	return false, time.Time{}, nil
}

func AddGiftBound(guildID, userID string, expiration time.Time) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := guildID + userID
	filter := bson.M{"_id": id}
	gb := giftBounds.FindOne(ctx, filter)

	if err = gb.Err(); err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			ctx, cancelOp := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancelOp()
			exp, err := expiration.UTC().GobEncode()
			if err != nil {
				return err
			}
			_, err = finance.InsertOne(ctx, bson.M{"_id": id, "gained": true, "expiration": exp})
			if err != nil {
				return err
			}
			return nil
		default:
			return err
		}
	}

	return nil
}
