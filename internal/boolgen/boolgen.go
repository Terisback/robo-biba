package boolgen

import (
	"math/rand"
	"time"
)

type Boolgen struct {
	src       rand.Source
	cache     int64
	remaining int
}

func New() *Boolgen {
	return &Boolgen{src: rand.NewSource(time.Now().UnixNano())}
}

func NewWithSrc(src rand.Source) *Boolgen {
	return &Boolgen{src: src}
}

func (bg *Boolgen) RandBool() bool {
	if bg.remaining == 0 {
		bg.cache, bg.remaining = bg.src.Int63(), 63
	}

	result := bg.cache&0x01 == 1
	bg.cache >>= 1
	bg.remaining--

	return result
}
