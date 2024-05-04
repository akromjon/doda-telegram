package app

import (
	"fmt"
	"os"
)

func DD(anything interface{}){
	fmt.Println(anything)
	os.Exit(0)
}

func Echo(value interface{}){
	fmt.Println(value)
}

func Format(format string, anything interface{}) string {
	return fmt.Sprintf(format,anything)
}