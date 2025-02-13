package decrypt

func StringDecrypt(text string) string {

	decryptedString := " "

	for _, str := range text {
		asciiCode := int(str)
		character := string(asciiCode - 3)
		decryptedString += character
	}
	return decryptedString
}
