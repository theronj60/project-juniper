package main

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
)

var question = []*survey.Question{
	{
		Name:      "cipher",
		Prompt:    &survey.Input{Message: "Which cipher would you like to use?"},
		Validate:  survey.Required,
		Transform: survey.Title,
	},
}

func main() {
	answers := struct {
		Cipher string
	}{}

	err := survey.Ask(question, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%s name", answers.Cipher)
	// v1
	// prompt for input
	// print to file
	// load file
}
