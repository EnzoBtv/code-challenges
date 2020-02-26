package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"unicode/utf8"
)

// Given the string, check if it is a palindrome.

// Example

// For inputString = "aabaa", the output should be checkPalindrome(inputString) = true; For inputString = "abac", the output should be checkPalindrome(inputString) = false; For inputString = "a", the output should be checkPalindrome(inputString) = true.

func revertString(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}

func isPalindromic(value string) bool {
	for i := 0; i < len(value)/2; i++ {
		if value[i] != value[len(value)-i-1] {
			return false
		}
	}
	return true
}

// Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

// Example 1:

// Input: "babad"
// Output: "bab"
// Note: "aba" is also a valid answer.

func longestPalindrome(s string) string {
	if len(s) == 0 || len(s) == 1 {
		return s
	}

	longest := string(s[0])

	for i := 0; i <= len(s)-1; i++ {
		current := string(s[i])
		for j := i + 1; j < len(s); j++ {
			current += string(s[j])
			if isPalindromic(current) && len(current) > len(longest) {
				longest = current
			}
		}
	}

	return longest
}

func convert(s string, numRows int) string {
	matrix := make([][]string, 0)
	letter := 0
	i := 0
	for letter < len(s) {
		matrix = append(matrix, make([]string, 0))
		for lettersPushed := 0; lettersPushed < numRows; lettersPushed++ {
			if letter == len(s) {
				break
			}
			matrix[i] = append(matrix[i], string(s[letter]))
			letter++
		}
		if letter == len(s) {
			break
		}
		i++
		if numRows-2 > 0 {
			matrix = append(matrix, make([]string, numRows))

			for j := numRows - 2; j > 0; j-- {
				if letter == len(s) {
					break
				}
				matrix[i][j] = string(s[letter])
				letter++
			}
			i++
		}
	}

	fullString := ""

	for j := 0; j < numRows; j++ {
		for i = 0; i < len(matrix); i++ {
			if j < len(matrix[i]) {
				fullString += matrix[i][j]

			}
		}
	}

	return fullString
}

// Given a 32-bit signed integer, reverse digits of an integer.

// Example 1:

// Input: 123
// Output: 321

func reverse(x int) int {
	strInt := strconv.Itoa(x)
	regexpNegative := regexp.MustCompile("-")
	negative := regexpNegative.MatchString(strInt)
	strInt = regexpNegative.ReplaceAllLiteralString(strInt, "")
	r := []rune(strInt)
	fmt.Println(r)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	if negative {
		strInt = fmt.Sprintf("-%s", string(r))
	} else {
		strInt = fmt.Sprintf("%s", string(r))
	}

	integer, _ := strconv.Atoi(strInt)

	if math.Pow(2, 31) < float64(integer) || -math.Pow(2, 31) > float64(integer) {
		return 0
	}

	return integer
}

func main() {
	fmt.Println(reverse(-2147483648))
}
