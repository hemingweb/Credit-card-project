package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter credit car number")
	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")
		cardNumber, _ := reader.ReadString('\n')
		// convert CRLF to LF
		cardNumber = strings.Replace(cardNumber, "\n", "", -1)

		isValid := Validation(cardNumber)

		if isValid {
			if response := CheckLuhn(cardNumber); response {
				fmt.Println("You entered a valid credit card number")
				os.Exit(0) // if valid exit
			} else {
				fmt.Println("You entered an invalid credit card number")
			}
		}
	}
}

func Validation(creditCardNumber string) bool {

	isValid := true
	// validate input length - if not 16 digits - exit
	if len(creditCardNumber) != 16 {
		isValid = false
		fmt.Println("Invalid credit card number length")
	}

	// validate characters - only digits allowed
	// isNotDigit := func(c rune) bool { return c < '0' || c > '9' }
	// isValid = strings.IndexFunc(creditCardNumber, isNotDigit) == -1
	// if !isValid {
	// 	fmt.Println("Invalid credit card characters")
	// }

	_, err := strconv.Atoi(creditCardNumber)
	if err != nil {
		isValid = false
		fmt.Println("Invalid credit card characters")
	}

	return isValid
}

func CheckLuhn(creditCardNumber string) bool {

	sum := 0
	isSecond := false

	for i := len(creditCardNumber) - 1; i >= 0; i-- {

		d := int(creditCardNumber[i] - '0')

		if isSecond {
			d = d * 2
		}

		sum += d / 10
		sum += d % 10

		isSecond = !isSecond
	}

	return sum%10 == 0
}
