package main

import (
	"fmt"
	"time"
)

func showMessage(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Millisecond * 300)
	}
}

func main() {
	go showMessage("GO ROUTINES")
	showMessage("LEARN GO!")
}
