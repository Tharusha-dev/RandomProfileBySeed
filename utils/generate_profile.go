package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func getFirstNameFile(gender int, firstNameDeterminingLetter string) string {

	var firstNameFile string

	if gender == 1 {
		log.Println(strings.ToUpper(firstNameDeterminingLetter) + ".txt" + "  Female")
		firstNameFile = "data/by_region/US/names/first_names/by_gender/female_" + strings.ToUpper(firstNameDeterminingLetter) + ".txt"
		log.Println(firstNameFile)

	} else {
		log.Println(strings.ToUpper(firstNameDeterminingLetter) + ".txt" + "  Male")
		firstNameFile = "data/by_region/US/names/first_names/by_gender/male_" + strings.ToUpper(firstNameDeterminingLetter) + ".txt"
		log.Println(firstNameFile)

	}

	return firstNameFile

}

func determineGender(genderDeterminingLetter *int) int {

	if *genderDeterminingLetter >= 78 && *genderDeterminingLetter <= 90 || *genderDeterminingLetter >= 110 && *genderDeterminingLetter <= 122 || *genderDeterminingLetter >= 52 && *genderDeterminingLetter <= 57 {

		return 0

	} else {

		return 1
	}

}

func findFirstTwoNonNumericalCharacters(seed string, asciiPoint map[string]int) (string, string) {
	var letters []int

	for _, l := range seed[2:] {
		kl := int(l)

		if kl >= asciiPoint["asciiLowerCaseLetterStarts"] && kl <= asciiPoint["asciiLowerCaseLetterEnds"] || kl >= asciiPoint["asciiUpperCaseLetterStarts"] && kl <= asciiPoint["asciiUpperCaseLetterEnds"] {

			letters = append(letters, kl)
		}

	}

	return string(letters[0]), string(letters[1])

}

func findIndexFromLetter(letter byte, asciiPoint map[string]int) int {

	l := int(letter)

	if l >= asciiPoint["asciiLowerCaseLetterStarts"] && l <= asciiPoint["asciiLowerCaseLetterEnds"] {
		return l - 96
	} else if l >= asciiPoint["asciiUpperCaseLetterStarts"] && l <= asciiPoint["asciiUpperCaseLetterEnds"] {
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

func getFormattedString(index int, fname string, lname string, seed string, fileToRead string) string {

	var formattedString string

	formattedString = getLineAtIndex(index, fileToRead)

	formattedString = strings.Replace(formattedString, "<fname>", strings.ToLower(fname), -1)
	formattedString = strings.Replace(formattedString, "<lname>", strings.ToLower(lname), -1)
	formattedString = strings.Replace(formattedString, "<int1>", fmt.Sprintf("%d", seed[0]), -1)
	formattedString = strings.Replace(formattedString, "<int2>", fmt.Sprintf("%d", seed[1]), -1)
	formattedString = strings.Replace(formattedString, "<int3>", fmt.Sprintf("%d", seed[2]), -1)
	formattedString = strings.Replace(formattedString, "<int4>", fmt.Sprintf("%d", seed[3]), -1)
	formattedString = strings.Replace(formattedString, "<int5>", fmt.Sprintf("%d", seed[4]), -1)

	return formattedString

}
