package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {

	var password string
	fmt.Println("Enter the Password ")
	fmt.Scanln(&password)

	fmt.Println("Password encrypted is", encShadow(password))

}

//creating a function to echeck md5 algo

func encShadow(password string) string {

	var hashcode = md5.Sum([]byte(password))
	return hex.EncodeToString(hashcode[:])
}
