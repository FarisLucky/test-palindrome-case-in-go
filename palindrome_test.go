package palindrome

import (
	"testing"
)

func TestIsNmbrPalindrome(t *testing.T) {
	var palindromeNmbr int = 1234321
	println(IsNmbrPalindrome(palindromeNmbr))
}

func TestIsStrPalindrome(t *testing.T) {
	var textPalindrome string = "amal lama"
	println(IsStrPalindrome(textPalindrome))
}
