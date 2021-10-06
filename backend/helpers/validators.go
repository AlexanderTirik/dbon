package helpers

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var validName = regexp.MustCompile("^[a-zA-Z0-9_]*$")
var validColName = regexp.MustCompile("^[a-zA-Z0-9_]*$")
var txtFileRegex = regexp.MustCompile(".txt$")
var intervals = [][]int{{0, 256}, {0, 256}, {0, 256}}

func ValidateName(input string) error {
	isNameValidErr := validName.FindStringSubmatch(input)
	if isNameValidErr == nil || len(input) == 0 {
		return errors.New("Invalid name")
	}
	return nil
}

func validateColName(input string) error {
	isNameValidErr := validName.FindStringSubmatch(input)
	if isNameValidErr == nil || input == "id" || len(input) == 0 {
		return errors.New("Invalid column name")
	}
	return nil
}

func validateInteger(input string) error {
	_, err := strconv.Atoi(input)
	if err != nil {
		return errors.New("Invalid integer")
	}
	return nil
}

func validateReal(input string) error {
	_, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return errors.New("Invalid real")
	}
	return nil
}

func validateChar(input string) error {
	validLetter := regexp.MustCompile("^[a-zA-Z]*$")
	isNameValidErr := validLetter.FindStringSubmatch(input)
	if isNameValidErr == nil || len(input) != 1 {
		return errors.New("Invalid char")
	}
	return nil
}

func validateTxtFile(fileName string) error {
	isNameValidErr := txtFileRegex.FindStringSubmatch(fileName)
	if isNameValidErr == nil {
		return errors.New("Not txt")
	}
	return nil
}

func IsValidType(t string) bool {
	return t == "integer" || t == "string" || t == "char" || t == "real" || t == "color" || t == "colorInvl"
}

func validateColor(input string, intervals [][]int) error {
	rgb := strings.Split(input, ",")
	if len(rgb) != 3 {
		return errors.New("Invalid color")
	}
	for i, c := range rgb {
		v, err := strconv.Atoi(c)
		if err != nil || v < intervals[i][0] || v > intervals[i][1] {
			return errors.New("Invalid color")
		}
	}

	return nil
}
