package app

import (
	"crypto/rand"
	"math/big"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
)

func DetectKey() keys.Key {
	var pressedKey keys.Key
	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		pressedKey = key
		return true, nil
	})
	return pressedKey
}

func GenerateDiscount() int {
	var discChance int64 = 40 //chance untuk dapet diskon
	var maxDisc int64 = 30    //diskon maksimum
	var minDisc int64 = 10    //diskon minimum
	disc, _ := rand.Int(rand.Reader, big.NewInt(100))
	if disc.Int64()+1 < discChance {
		return int((disc.Int64()+1)*(maxDisc-minDisc))/100 + int(minDisc)
	} else {
		return 0
	}
}
