package main

import (
	"github.com/sjkaliski/notify"
)

func main() {
	notify.Alert(&notify.Message{
		Title: "Foo",
		Body:  "Bar",
	})
}
