package caeser

import "fmt"

// read from text file
// should encrypt or decrypt text
// convert lines to runes?
// take a single letter and shift it right 3 spots in the alphabet to encrpyt
// reverse it with decrypting

// create a variable (array) to convert spaces to a random but predefined character
// var spaceChar string = {}array

func readFile() {
	//
}

func writeFile() {
	//
}

func Encrypt(str string, method string) string {
	if method == "file" {
		// for i, r := range str {
		// 	fmt.Println("encrypt caeser file")
		// }
		// needs to read file
		// then loop through string
	} else {
		var encryptedTxt []string
		for _, r := range str {
			encryptedTxt = append(encryptedTxt, string(r + 3))
		} 
		fmt.Println(encryptedTxt)
		// needs to loop through string
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
