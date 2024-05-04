package app

import "os"

type Command struct {
	command string
	value   string
}

func Capture() Command {

	if len(os.Args) < 3 {

		Help()

	}

	if os.Args[1] == "help" {

		Help()
	}

	return Command{
		command: os.Args[1],
		value:   os.Args[2],
	}
}
