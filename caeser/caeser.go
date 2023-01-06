package caeser

import (
	"fmt"
	"math/rand"
	"strings"
)

// read from text file
// should encrypt or decrypt text
// convert lines to runes?
// take a single letter and shift it right 3 spots in the alphabet to encrpyt
// reverse it with decrypting

// create a variable (array) to convert spaces to a random but predefined character
var spaceChar = [...]string{"&","$","*","#"}

func readFile() {
	//
}

func writeFile() {
	//
}

func Encrypt(str string, method string) string {
	var encrpytedValue string
	if method == "file" {
		// for i, r := range str {
		// 	fmt.Println("encrypt caeser file")
		// }
		// needs to read file
		// then loop through string
	} else {
		var encryptedArray []string
		for _, r := range str {
			if r == 32 {
				randIndex := rand.Intn(len(spaceChar))
				encryptedArray = append(encryptedArray, string(spaceChar[randIndex]))
			} else {
				encryptedArray = append(encryptedArray, string(r + 3))
			}
		} 
		encrpytedValue = strings.Join(encryptedArray, "")
	}
	return encrpytedValue
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
