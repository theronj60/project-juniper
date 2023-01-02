package main

import (
	"fmt"
	"flag"
	"github.com/AlecAivazis/survey/v2"
	"github.com/theronj60/project-juniper/caeser"
)

// var question = []*survey.Question{
// 	{
// 		Name:      "cipher",
// 		Prompt:    &survey.Input{Message: "Which cipher would you like to use?"},
// 		Validate:  survey.Required,
// 		Transform: survey.Title,
// 	},
// }

var question = &survey.Select{
	Message: "Which cipher would you like to use?",
	Options: []string{"caeser", "vigenere"},
}

var prompt = &survey.Input{
	Message: "Enter your value to encrypt.",
}

func main() {
	var cipher string = ""
	var value string = ""
	file := flag.Bool("f", false, "a bool")
	text := flag.Bool("t", false, "a bool")

	flag.Parse()

	err := survey.AskOne(question, &cipher)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	survey.AskOne(prompt, &value)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// @TODO need to add an argument for specifying encrypt or decrypt

	switch cipher {
	case "caeser":
		if *file {
			caeser.Encrypt(value, "file")
			fmt.Println("file is true")
		}
		if *text {
			caeser.Encrypt(value, "text")
			fmt.Println("text is true")
		}	
	case "vigenere":
		fmt.Println("vigenere")
	default:
		fmt.Println("default")
	}
}
