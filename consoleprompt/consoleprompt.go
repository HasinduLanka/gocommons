package consoleprompt

import (
	"fmt"
	"strings"

	"github.com/HasinduLanka/gocommons/console"
)

var NoConsole bool = false

func ReadLine() string {
	var s string
	if NoConsole {
		s = ""
	} else {
		fmt.Scanln(&s)
	}
	return s
}

func Prompt(msg string) string {
	console.Print(msg)
	return ReadLine()
}

func PromptOptions(msg string, options map[string]string) string {
	console.Print(msg)
	for o, m := range options {
		console.Print("\t[" + o + "] = " + m)
	}

	var r string = ""
	if NoConsole {
		// Select First key
		for o := range options {
			r = o
			break
		}
	} else {
		r = strings.TrimSpace(strings.ToLower(Prompt("Enter [value] : ")))
	}

	_, ok := options[r]
	if ok {
		return r
	} else {
		console.Print("Sorry, I didn't get that. Please enter the [option] you want ")
		return PromptOptions(msg, options)
	}

}
