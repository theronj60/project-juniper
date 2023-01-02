package main

import (
	"fmt"
	"strings"
	"github.com/AlecAivazis/survey/v2"
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

	value = strings.ToLower(value)

	fmt.Println("Cipher: ", cipher, "\nValue: ", value)
	// v1
	// prompt for input
	// print to file
	// load file
}
