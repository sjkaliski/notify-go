package main

import (
	"github.com/sjkaliski/notify-go"
)

func main() {
	notify.Alert(&notify.Message{
		Title: "Foo",
		Body:  "Bar",
	})
}
