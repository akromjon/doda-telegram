package app

import "errors"

func StartTheApp() (bool, error) {

	config, _ := Boot()

	telegram := NewTelegram(&config)

	command := Capture()

	if command.command == "message" {

		t, err := telegram.SendMessage(&command.value)

		return t, err

	} else if command.command == "file" {

		t, err := telegram.SendFile(&command.value)

		return t, err
	}

	msg := command.command

	return false, errors.New(msg + " - is not supported!")

}
