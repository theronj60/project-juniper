package main

import (
	"flag"
	"fmt"
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

var promptEncrypt = &survey.Input{
	Message: "Enter your value to encrypt.",
}

var promptDecrypt = &survey.Input{
	Message: "Enter your value to decrypt.",
}

func main() {
	var cipher string = ""
	var value string = ""
	file := flag.Bool("f", false, "sets cipher to use a file and parse it.")
	text := flag.Bool("t", false, "sets cipher to use provided string.")
	encrypt := flag.Bool("encrypt", false, "sets cipher to encrypt.")
	decrypt := flag.Bool("decrypt", false, "sets cipher to decrypt.")

	flag.Parse()

	err := survey.AskOne(question, &cipher)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	func(encrypt bool, decrypt bool) {
		if encrypt {
			survey.AskOne(promptEncrypt, &value)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}
		if decrypt {
			survey.AskOne(promptDecrypt, &value)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}
	}(*encrypt, *decrypt)

	// @TODO add error handling and hints

	switch cipher {
	case "caeser":
		if *file {
			if *encrypt {
				caeser.Encrypt(value, "file")
			}
			if *decrypt {
				caeser.Decrypt(value, "file")
			}
			fmt.Println("file is true")
		}
		if *text {
			if *encrypt {
				// @TODO see if we can do a return instead
				fmt.Println(caeser.Encrypt(value, "text"))
			}
			if *decrypt {
				fmt.Println(caeser.Decrypt(value, "text"))
			}
		}
	case "vigenere":
		fmt.Println("vigenere")
	default:
		fmt.Println("default")
	}
}
