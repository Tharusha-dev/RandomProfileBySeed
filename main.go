package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	const userNameFile = "usernames.txt"
	const emailFormatsFile = "emails.txt"
	const usAddressesFile = "addresses/addresses.txt"

	// File
	var fnameFile string

	// stdio seed enter
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter seed:")

	read, err_read := reader.ReadString('\n')
	if err_read != nil {
		log.Println(err_read)
	}

	//gender letter is set to the SECOND letter
	gender_letter := int(read[1])

	//fname letter is the first non numerical letter and lname letter is the second non numerical letter
	fname_letter, lname_letter := findFirstTwoNonNumericalCharacters(read)

	// 0 => Male & 1 => Female
	gender := getGender(&gender_letter)

	if gender == 1 {
		log.Println(strings.ToUpper(fname_letter) + ".txt" + "  Female")
		fnameFile = "by_gender/female_" + strings.ToUpper(fname_letter) + ".txt"
		log.Println(fnameFile)

	} else {
		log.Println(strings.ToUpper(fname_letter) + ".txt" + "  Male")
		fnameFile = "by_gender/male_" + strings.ToUpper(fname_letter) + ".txt"
		log.Println(fnameFile)

	}

	// first name offset is calculated by THIRD letter * FOURTH letter
	fname_offset := findIndexFromLetter(read[2]) * findIndexFromLetter(read[3])

	// last name offset is calculated by FOUTH letter * FIFTH (LAST) letter
	lname_offset := findIndexFromLetter(read[3]) * findIndexFromLetter(read[4])

	// address offset is calculated by SECOND letter * FOURTH letter
	address_offset := findIndexFromLetter(read[1]) * findIndexFromLetter(read[3])

	// email offset set to the ascii of FORTH letter
	email_offset := int(read[3])

	// username offset set to the ascii of FIFTH (LAST) letter
	username_offset := int(read[4])

	fname := getLineAtIndex(fname_offset, fnameFile)

	lnameFile := "last_names/lname_" + strings.ToUpper(lname_letter) + ".txt"

	lname := getLineAtIndex(lname_offset, lnameFile)

	email := getEmail(email_offset, fname, lname, read, emailFormatsFile)

	address := getLineAtIndex(address_offset, usAddressesFile)

	username := getUsername(username_offset, fname, lname, read, userNameFile)

	log.Println("test ****")
	log.Println(fname)
	log.Println(lname)
	log.Println(email)
	log.Println(address)
	log.Println(username)

}

func getGender(secondLetter *int) int {

	if *secondLetter >= 78 && *secondLetter <= 90 || *secondLetter >= 110 && *secondLetter <= 122 || *secondLetter >= 52 && *secondLetter <= 57 {

		return 0

	} else {

		return 1
	}

}

func findFirstTwoNonNumericalCharacters(seed string) (string, string) {
	var letters []int

	for _, l := range seed[2:] {
		kl := int(l)

		if kl >= 97 && kl <= 122 || kl >= 65 && kl <= 90 {

			letters = append(letters, kl)
		}

	}

	return string(letters[0]), string(letters[1])

}

func findIndexFromLetter(letter byte) int {

	l := int(letter)

	if l >= 97 && l <= 122 {
		return l - 96
	} else if l >= 65 && l <= 90 {
		return l - 38
	} else {
		return l + 5
	}

}

func getLineAtIndex(index int, fileToRead string) string {

	linesread := 0
	var line string

	for {

		if linesread >= index-1 {

			break
		}

		file, err_file := os.Open(fileToRead)

		if err_file != nil {
			log.Printf("could not open the file: %v", err_file)
		}

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {

			line = scanner.Text()

			if linesread >= index-1 {

				break
			}

			linesread++

		}

		file.Close()

	}

	return line
}

func getEmail(index int, fname string, lname string, seed string, fileToRead string) string {

	var email string

	email = getLineAtIndex(index, fileToRead)
	email = getStringFromTemplate(&email, seed, fname, lname)

	return email

}

func getUsername(index int, fname string, lname string, seed string, fileToRead string) string {

	var username string

	username = getLineAtIndex(index, fileToRead)
	username = getStringFromTemplate(&username, seed, fname, lname)

	return username

}

func getStringFromTemplate(original *string, seed string, fname string, lname string) string {

	*original = strings.Replace(*original, "<fname>", strings.ToLower(fname), -1)
	*original = strings.Replace(*original, "<lname>", strings.ToLower(lname), -1)
	*original = strings.Replace(*original, "<int1>", fmt.Sprintf("%d", seed[0]), -1)
	*original = strings.Replace(*original, "<int2>", fmt.Sprintf("%d", seed[1]), -1)
	*original = strings.Replace(*original, "<int3>", fmt.Sprintf("%d", seed[2]), -1)
	*original = strings.Replace(*original, "<int4>", fmt.Sprintf("%d", seed[3]), -1)
	*original = strings.Replace(*original, "<int5>", fmt.Sprintf("%d", seed[4]), -1)

	return *original

}
