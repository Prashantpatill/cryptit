package main

import (
	"fmt"

	"github.com/Prashantpatill/cryptit/decrypt"
	"github.com/Prashantpatill/cryptit/encrypt"
)

func main() {
	var input string
	fmt.Println("Enter the String you want to encrypt")
	fmt.Scanf("%s", &input)
	fmt.Println("Begining to encrypt")
	encryptText := encrypt.StringEncryption(input)
	fmt.Println("Encrypted Text is", encryptText)
	fmt.Println("Ecnryptiion Completed")
	fmt.Println("Begining to decrypt")
	decryptText := decrypt.StringDecrypt(encryptText)
	fmt.Println("Decrypted Text is", decryptText)
	fmt.Println("Decrytpion Completed")
}
