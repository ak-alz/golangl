package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	strings := [3]string{"a4bc2d5e", "abcd", "45"}

	for _, v := range strings {
		newString, err := stringPrepare(v)

		if err != nil {
			fmt.Printf("string %s is incorrect!!!\n", v)
		} else {
			fmt.Printf("New string is: \"%s\"\n", newString)
		}
	}
}

func stringPrepare(str string) (newString string, err error) {
	var previous string

	for _, char := range str {
		strChar := string(char)
		isNumeric := isNumeric(strChar)

		if !isNumeric {
			previous = strChar
			newString += previous
			continue
		}

		if previous == "" {
			err = errors.New(fmt.Sprint("Invalid input string"))
		}

		i, _ := strconv.Atoi(strChar)
		for iter := 1; iter < i; iter++ {
			newString += previous
		}
	}
	return
}

func isNumeric(char string) bool {
	_, err := strconv.ParseFloat(char, 64)
	return err == nil
}
