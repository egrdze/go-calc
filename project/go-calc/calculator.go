package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var input1, input2 string
	var operation string

	fmt.Print("Enter the first number (Arabic or Roman numeral): ")
	fmt.Scanln(&input1)

	fmt.Print("Enter the operation (+, -, *, /): ")
	fmt.Scanln(&operation)

	fmt.Print("Enter the second number (Arabic or Roman numeral): ")
	fmt.Scanln(&input2)

	isRoman1 := isRomanNumeral(input1)
	isRoman2 := isRomanNumeral(input2)

	if isRoman1 != isRoman2 {
		fmt.Println("Error: Mixing Roman and Arabic numerals is not allowed.")
		return
	}

	var result string
	if isRoman1 {
		arabic1 := romanToArabic(input1)
		arabic2 := romanToArabic(input2)

		switch operation {
		case "+":
			result = arabicToRoman(arabic1 + arabic2)
		case "-":
			result = arabicToRoman(arabic1 - arabic2)
		case "*":
			result = arabicToRoman(arabic1 * arabic2)
		case "/":
			if arabic2 == 0 {
				fmt.Println("Error: Division by zero is not allowed.")
				return
			}
			result = arabicToRoman(arabic1 / arabic2)
		default:
			fmt.Println("Error: Invalid operation.")
			return
		}
	} else {
		arabic1, err1 := strconv.Atoi(input1)
		arabic2, err2 := strconv.Atoi(input2)

		if err1 != nil || err2 != nil {
			fmt.Println("Error: Invalid Arabic numerals.")
			return
		}

		switch operation {
		case "+":
			result = strconv.Itoa(arabic1 + arabic2)
		case "-":
			result = strconv.Itoa(arabic1 - arabic2)
		case "*":
			result = strconv.Itoa(arabic1 * arabic2)
		case "/":
			if arabic2 == 0 {
				fmt.Println("Error: Division by zero is not allowed.")
				return
			}
			result = strconv.Itoa(arabic1 / arabic2)
		default:
			fmt.Println("Error: Invalid operation.")
			return
		}
	}

	fmt.Printf("Result: %s\n", result)
}

func isRomanNumeral(input string) bool {
	validChars := "IVXLCDM"
	for _, char := range input {
		if !strings.ContainsRune(validChars, char) {
			return false
		}
	}
	return true
}

func romanToArabic(input string) int {
	return toArabic(input)
}

func arabicToRoman(arabic int) string {
	if arabic <= 0 {
		return "Invalid"
	}

	romanSymbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	romanValues := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	roman := ""

	for i := 0; i < len(romanSymbols); i++ {
		for arabic >= romanValues[i] {
			roman += romanSymbols[i]
			arabic -= romanValues[i]
		}
	}

	return roman
}

func toArabic(input string) int {
	arabic, err := strconv.Atoi(input)
	if err == nil {
		return arabic
	}

	romanValues := map[string]int{
		"M": 1000,
		"D": 500,
		"C": 100,
		"L": 50,
		"X": 10,
		"V": 5,
		"I": 1,
	}

	arabicValue := 0
	prevValue := 0

	for i := len(input) - 1; i >= 0; i-- {
		char := string(input[i])
		value, exists := romanValues[char]

		if !exists {
			return 0 
		}

		if value >= prevValue {
			arabicValue += value
		} else {
			arabicValue -= value
		}

		prevValue = value
	}

	return arabicValue
}
