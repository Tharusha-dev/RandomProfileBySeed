package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var fileToRead string

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter :")

	read, err_read := reader.ReadString('\n')
	if err_read != nil {
		fmt.Println(err_read)
	}

	gender_letter := int(read[1])

	fname_letter, _ := getStrings(read)

	gender := getGender(&gender_letter) // 0 => Male & 1 => Female

	if gender == 1 {
		fmt.Println(strings.ToUpper(fname_letter) + ".txt" + "  Female")
		fileToRead = "by_gender/female_" + strings.ToUpper(fname_letter) + ".txt"
		fmt.Println(fileToRead)

	} else {
		fmt.Println(strings.ToUpper(fname_letter) + ".txt" + "  Male")
		fileToRead = "by_gender/male_" + strings.ToUpper(fname_letter) + ".txt"
		fmt.Println(fileToRead)

	}

	fname_offset := findIndexFromLetters(read[2]) * findIndexFromLetters(read[3])

	nameTest := getNameAtIndex(&fname_offset, &fileToRead)

	fmt.Println(fname_offset)

	fmt.Println(gender)

	fmt.Println("test ****")
	fmt.Println(nameTest)

}

func prepRanges() ([]int, []int) {

	var lowerCase []int
	var upperCase []int

	// lowerCase = make([]int, 26,26)

	for i := 97; i <= 122; i++ {
		lowerCase = append(lowerCase, i)

	}

	for i := 65; i <= 90; i++ {
		upperCase = append(upperCase, i)

	}

	return lowerCase, upperCase
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
			fmt.Println("initial scan linesread:- " + fmt.Sprint(linesread))

			name = scanner.Text()

			if linesread >= *index-1 {
				fmt.Println("broke")

				break
			}

			linesread++

		}

		file.Close()

	}

	return name
}
