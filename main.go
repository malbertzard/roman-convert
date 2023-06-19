package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var romanValues = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func isStringRoman(str string) bool {
	for i := 0; i < len(str); i++ {
		if _, ok := romanValues[str[i]]; !ok {
			return false
		}
	}
	return true
}

func romanToInt(roman string) int {
	result := 0
	prevValue := 0
	value := 0

	for i := len(roman) - 1; i >= 0; i-- {
		value = romanValues[roman[i]]
		if value < prevValue {
			result -= value
		} else {
			result += value
			prevValue = value
		}
	}

	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter a Roman numeral (or 'q' to quit): ")
		roman, _ := reader.ReadString('\n')
		roman = strings.TrimSuffix(roman, "\n")
		roman = strings.ToUpper(roman)

		if roman == "Q" {
			break
		}

		if isStringRoman(roman) {
			result := romanToInt(roman)
			fmt.Println("The equivalent integer value is:", result)
		} else {
			fmt.Println("The input is not a Roman numeral")
		}
	}
}
