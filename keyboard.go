package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

func Interrupt() bool {
	if err := keyboard.Open(); err != nil {
		Exit(err.Error())
	}
	defer func() {
		err := keyboard.Close()
		if err != nil {
			Exit(err.Error())
		}
	}()

	fmt.Println("Press any key to start the quiz. Press ESC to quit.")

	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			Exit(err.Error())
		}

		if key != keyboard.KeyEsc {
			return false
		}
		return true
	}
}
