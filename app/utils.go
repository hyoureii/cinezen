package app

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
)

func DetectKey() keys.Key {
	var pressedKey keys.Key
	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		if key.Code == keys.CtrlX { //exit aplikasi kapanpun ketika user menekan ctrl + x
			fmt.Println("\n\rKeluar aplikasi")
			os.Exit(0)
		}
		pressedKey = key
		return true, nil
	})
	return pressedKey
}

func InputText() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	return text
}
