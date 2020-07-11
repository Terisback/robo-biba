package utils

import (
	"math/rand"
	"time"
)

type boolgen struct {
	src       rand.Source
	cache     int64
	remaining int
}

var (
	bg = &boolgen{src: rand.NewSource(time.Now().UnixNano())}
)

func RandBool() bool {
	if bg.remaining == 0 {
		bg.cache, bg.remaining = bg.src.Int63(), 63
	}

	result := bg.cache&0x01 == 1
	bg.cache >>= 1
	bg.remaining--

	return result
}
