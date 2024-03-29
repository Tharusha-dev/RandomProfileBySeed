package utils

import (
	"fmt"
	"strings"
)

func GetProfileFromSeed(seed string) (string, string, string, error) {

	asciiPoints := make(map[string]int)
	asciiPoints["asciiLowerCaseLetterStarts"] = 97
	asciiPoints["asciiLowerCaseLetterEnds"] = 122
	asciiPoints["asciiUpperCaseLetterStarts"] = 65
	asciiPoints["asciiUpperCaseLetterEnds"] = 90

	fmt.Println(asciiPoints)

	const userNameFile = "data/common_templates/usernames.txt"
	const emailFormatsFile = "data/common_templates/emails.txt"
	const usAddressesFile = "data/by_region/US/addresses/addresses.txt"

	// seed := getInput()

	//gender letter is set to the SECOND letter
	genderDeterminingLetter := int(seed[1])

	//fname letter is the first non numerical letter and lname letter is the second non numerical letter
	firstNameDeterminingLetter, lastNameDeterminingLetter := findFirstTwoNonNumericalCharacters(seed, asciiPoints)

	// 0 => Male & 1 => Female
	gender := determineGender(&genderDeterminingLetter)

	firstNameFile := getFirstNameFile(gender, firstNameDeterminingLetter)

	firstNameOffset, lastNameOffset, addressOffset, emailOffset, usernameOffset := getOffsets(seed, asciiPoints)

	firstName := getLineAtIndex(firstNameOffset, firstNameFile)

	lastNameFile := "data/by_region/US/names/last_names/lname_" + strings.ToUpper(lastNameDeterminingLetter) + ".txt"

	lastName := getLineAtIndex(lastNameOffset, lastNameFile)

	email := getFormattedString(emailOffset, firstName, lastName, seed, emailFormatsFile)

	address := getLineAtIndex(addressOffset, usAddressesFile)

	username := getFormattedString(usernameOffset, firstName, lastName, seed, userNameFile)

	return email, address, username, nil

	// generatedProfile := Profile{firstName: firstName, lastName: lastName, username: username, email: email, address: address}

}
