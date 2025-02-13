package encrypt

func StringEncryption(text string) string {

	encryptedString := " "

	for _, str := range text {
		asciiCode := int(str)
		character := string(asciiCode + 3)
		encryptedString += character
	}
	return encryptedString
}
