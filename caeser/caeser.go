package caeser

import "fmt"

// read from text file
// should encrypt or decrypt text
// convert lines to runes

func readFile() {
	//
}

func writeFile() {
	//
}

func Encrypt(str string, method string) string {
	if method == "file" {
		fmt.Println("encrypt caeser file")
		//
	} else {
		//
		fmt.Println("encrypt caeser text")
	}
	return "string encrypt"
}

func Decrypt(str string, method string) string {
	if method == "file" {
		fmt.Println("decrypt caeser file")
		//
	} else {
		fmt.Println("decrypt caeser text")
		//
	}
	return "string decrypt"
}
