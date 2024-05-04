package main

import (
	"github.com/akromjon/doda-telegram-cli/app"
)

func main() {

	_, err := app.StartTheApp()

	if err != nil {

		app.DD(err)

	}

}
