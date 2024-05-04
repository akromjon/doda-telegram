package app

import (
	"errors"
	"os"
)

type Command struct {
	command string
	value   string
}

func Capture() (Command, error) {

	if len(os.Args) < 3 {

		Help()

	}

	if os.Args[1] == "help" {

		Help()
	}

	if os.Args[1] == "message" {

		return Command{
			command: "message",
			value:   os.Args[2],
		}, nil
	} else if os.Args[1] == "file" {

		return Command{
			command: "file",
			value:   os.Args[2],
		}, nil
	}

	return Command{}, errors.New("The command is not supported")

}
