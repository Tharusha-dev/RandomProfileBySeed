package utils

func getOffsets(seed string, asciiPoints map[string]int) (int, int, int, int, int) {

	// first name offset is calculated by THIRD letter * FOURTH letter
	firstNameOffset := findIndexFromLetter(seed[2], asciiPoints) * findIndexFromLetter(seed[3], asciiPoints)

	// last name offset is calculated by FOUTH letter * FIFTH (LAST) letter
	lastNameOffset := findIndexFromLetter(seed[3], asciiPoints) * findIndexFromLetter(seed[4], asciiPoints)

	// address offset is calculated by SECOND letter * FOURTH letter
	addressOffset := findIndexFromLetter(seed[1], asciiPoints) * findIndexFromLetter(seed[3], asciiPoints)

	// email offset set to the ascii of FORTH letter
	emailOffset := int(seed[3])

	// username offset set to the ascii of FIFTH (LAST) letter
	usernameOffset := int(seed[4])

	return firstNameOffset, lastNameOffset, addressOffset, emailOffset, usernameOffset

}
