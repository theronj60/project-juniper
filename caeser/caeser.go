package caeser

import (
	"fmt"
	"math/rand"
	"strings"
)

// read from text file
// should encrypt or decrypt text
// take a single letter and shift it right 3 spots in the alphabet to encrpyt

// create a variable (array) to convert spaces to a random but predefined character
// needed? no, good start for something? maybe
var spaceChar = []string{"&","$","*","#"}

// helper functions
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
// end helper functions

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
	var decryptedValue string

	if method == "file" {
		fmt.Println("decrypt caeser file")
	} else {
		var decryptedArray []string
		for _, r := range str {
			if contains(spaceChar, string(r)) {
				decryptedArray = append(decryptedArray, " ")
			} else {
				decryptedArray = append(decryptedArray, string(r - 3))
			}
		}
		decryptedValue = strings.Join(decryptedArray, "")
	}
	return decryptedValue
}
