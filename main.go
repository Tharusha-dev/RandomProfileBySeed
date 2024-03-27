package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// Ffile
	var fileToRead string

	// stdio seed enter
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter :")

	read, err_read := reader.ReadString('\n')
	if err_read != nil {
		fmt.Println(err_read)
	}

	//gender letter is set to the SECOND letter
	gender_letter := int(read[1])

	//fname letter is the first non numerical letter and lname letter is the second non numerical letter
	fname_letter, lname_letter := getStrings(read)

	// 0 => Male & 1 => Female
	gender := getGender(&gender_letter)

	if gender == 1 {
		fmt.Println(strings.ToUpper(fname_letter) + ".txt" + "  Female")
		fileToRead = "by_gender/female_" + strings.ToUpper(fname_letter) + ".txt"
		fmt.Println(fileToRead)

	} else {
		fmt.Println(strings.ToUpper(fname_letter) + ".txt" + "  Male")
		fileToRead = "by_gender/male_" + strings.ToUpper(fname_letter) + ".txt"
		fmt.Println(fileToRead)

	}

	// first name offset is calculated by THIRD letter * FOURTH letter
	fname_offset := findIndexFromLetters(read[2]) * findIndexFromLetters(read[3])

	// last name offset is calculated by FOUTH letter * FIFTH (LAST) letter
	lname_offset := findIndexFromLetters(read[3]) * findIndexFromLetters(read[4])

	// address offset is calculated by SECOND letter * FOURTH letter
	address_offset := findIndexFromLetters(read[1]) * findIndexFromLetters(read[3])

	// email offset set to the ascii of FORTH letter
	email_offset := int(read[3])

	// username offset set to the ascii of FIFTH (LAST) letter
	username_offset := int(read[4])

	fname := getNameAtIndex(&fname_offset, &fileToRead)

	fileToRead = "last_names/lname_" + strings.ToUpper(lname_letter) + ".txt"

	lname := getNameAtIndex(&lname_offset, &fileToRead)

	fileToRead = "emails.txt"

	email := getEmail(email_offset, fname, lname, read, &fileToRead)

	fileToRead = "addresses/addresses.txt"

	address := getNameAtIndex(&address_offset, &fileToRead)

	fileToRead = "usernames.txt"

	username := getUsername(username_offset, fname, lname, read, &fileToRead)

	fmt.Println(fname_offset)

	fmt.Println(gender)

	fmt.Println("test ****")
	fmt.Println(fname)
	fmt.Println(lname)
	fmt.Println(email)
	fmt.Println(address)
	fmt.Println(username)

}

func getGender(secondLetter *int) int {

	if *secondLetter >= 78 && *secondLetter <= 90 || *secondLetter >= 110 && *secondLetter <= 122 || *secondLetter >= 52 && *secondLetter <= 57 {

		return 0

	} else {

		return 1
	}

}

func getStrings(seed string) (string, string) {
	var letters []int

	for _, l := range seed[2:] {
		kl := int(l)

		if kl >= 97 && kl <= 122 || kl >= 65 && kl <= 90 {

			letters = append(letters, kl)
		}

	}

	return string(letters[0]), string(letters[1])

}

func findIndexFromLetters(letter1 byte) int {

	l := int(letter1)

	if l >= 97 && l <= 122 {
		return l - 96
	} else if l >= 65 && l <= 90 {
		return l - 38
	} else {
		return l + 5
	}

}

func getNameAtIndex(index *int, fileToRead *string) string {

	linesread := 0
	var name string

	for {

		if linesread >= *index-1 {

			break
		}

		file, err_file := os.Open(*fileToRead)

		if err_file != nil {
			fmt.Println("could not open the file: %v", err_file)
		}

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			// fmt.Println("initial scan linesread:- " + fmt.Sprint(linesread))

			name = scanner.Text()

			if linesread >= *index-1 {

				break
			}

			linesread++

		}

		file.Close()

	}

	return name
}

func getEmail(index int, fname string, lname string, seed string, fileToRead *string) string {

	linesread := 0
	var email string

	for {

		if linesread >= index-1 {

			break
		}

		file, err_file := os.Open(*fileToRead)

		if err_file != nil {
			fmt.Println("could not open the file: %v", err_file)
		}

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			// fmt.Println("initial scan linesread:- " + fmt.Sprint(linesread))

			email = scanner.Text()

			if linesread >= index-1 {

				break
			}

			linesread++

		}

		file.Close()

	}

	email = strings.Replace(email, "<fname>", strings.ToLower(fname), -1)
	email = strings.Replace(email, "<lname>", strings.ToLower(lname), -1)
	email = strings.Replace(email, "<int1>", fmt.Sprintf("%d", seed[0]), -1)
	email = strings.Replace(email, "<int2>", fmt.Sprintf("%d", seed[1]), -1)
	email = strings.Replace(email, "<int3>", fmt.Sprintf("%d", seed[2]), -1)
	email = strings.Replace(email, "<int4>", fmt.Sprintf("%d", seed[3]), -1)
	email = strings.Replace(email, "<int5>", fmt.Sprintf("%d", seed[4]), -1)

	return email

}

func getUsername(index int, fname string, lname string, seed string, fileToRead *string) string {

	linesread := 0
	var username string

	for {

		if linesread >= index-1 {

			break
		}

		file, err_file := os.Open(*fileToRead)

		if err_file != nil {
			fmt.Println("could not open the file: %v", err_file)
		}

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			// fmt.Println("initial scan linesread:- " + fmt.Sprint(linesread))

			username = scanner.Text()

			if linesread >= index-1 {

				break
			}

			linesread++

		}

		file.Close()

	}

	username = strings.Replace(username, "<fname>", strings.ToLower(fname), -1)
	username = strings.Replace(username, "<lname>", strings.ToLower(lname), -1)
	username = strings.Replace(username, "<int1>", fmt.Sprintf("%d", seed[0]), -1)
	username = strings.Replace(username, "<int2>", fmt.Sprintf("%d", seed[1]), -1)
	username = strings.Replace(username, "<int3>", fmt.Sprintf("%d", seed[2]), -1)
	username = strings.Replace(username, "<int4>", fmt.Sprintf("%d", seed[3]), -1)
	username = strings.Replace(username, "<int5>", fmt.Sprintf("%d", seed[4]), -1)

	return username

}
