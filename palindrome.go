package palindrome

import "fmt"

func IsStrPalindrome(name string) string {
	if name == "" {
		return "input is empty"
	}
	// input := []rune(name)
	size := len(name) - 1
	for i := 0; i < size; i++ {
		if name[i] != name[size-i] {
			return "isn't palindrome"
		}
	}
	return "is palindrome"
}

func IsNmbrPalindrome(input int) string {
	var remainder, temp int
	var reverse int = 0
	if input == 0 {
		return "not be zero"
	}
	temp = input
	for input != 0 {
		remainder = input % 10
		reverse = reverse*10 + remainder
		input /= 10
		fmt.Printf("==== process ==== \nremainder = %d \nreverse = %d \ninput = %d \n==== end process ====\n\n", remainder, reverse, input)
	}
	if temp == reverse {
		return "is palindrome"
	}
	return "not palindrome"
}
