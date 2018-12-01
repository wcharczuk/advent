package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	a = byte(97)
)

func test(actual, expected interface{}) {
	fmt.Printf("test actual: %v expected: %v", actual, expected)
	if actual != expected {
		fmt.Println(" failed!")
		os.Exit(1)
	} else {
		fmt.Println(" passed!")
	}
}

func main() {
	password := "hxbxxyzz"
	fmt.Println("Input Password:", password)
	newPassword := generateNewPassword(password)
	fmt.Println("Final Password:", newPassword)
}

func generateNewPassword(password string) string {
	newPassword := password
	success := false
	for !success {
		newPassword = incrementPassword(newPassword)
		//fmt.Printf("trying %s\n", newPassword)
		success = checkPassword(newPassword)
	}
	return newPassword
}

func checkPassword(password string) bool {
	return hasStraight(password) && hasRepeats(password) && doesNotContain(password, []string{"i", "o", "l"})
}

func incrementPassword(password string) string {
	output := []byte(password)
	shouldCarry := true
	incremented := byte(0)
	for x := len(password) - 1; x >= 0 && shouldCarry; x-- {
		incremented, shouldCarry = incrementLetter(output[x])
		output[x] = incremented
		if x == 0 && shouldCarry {
			output = append([]byte{a}, output...)
		}
	}
	return string(output)
}

func incrementLetter(input byte) (byte, bool) {
	output := input + byte(1)
	if output > 122 {
		output = a
		return output, true
	}

	return output, false
}

func hasStraight(input string) bool {
	for x := 0; x < len(input)-2; x++ {
		basis := input[x]
		basisNext, shouldCarryNext := incrementLetter(basis)
		if shouldCarryNext {
			continue
		}
		basisNextNext, shouldCarryNextNext := incrementLetter(basisNext)
		if shouldCarryNextNext {
			continue
		}

		if input[x+1] == basisNext && input[x+2] == basisNextNext {
			return true
		}
	}

	return false
}

func doesNotContain(input string, set []string) bool {
	for _, bad := range set {
		if strings.Contains(input, bad) {
			return false
		}
	}

	return true
}

func hasRepeats(input string) bool {
	repeatCount := 0
	for x := 0; x < len(input)-1; x++ {
		basis := input[x]
		basisNext := input[x+1]
		if basis == basisNext {
			repeatCount = repeatCount + 1
			x = x + 1
			if repeatCount == 2 {
				return true
			}
		}
	}

	return false
}
