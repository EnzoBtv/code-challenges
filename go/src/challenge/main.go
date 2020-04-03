package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
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

func wordBreak(input string, tokens []string) bool {
	size := len(input)
	if size == 0 {
		return true
	}

	tokenMap := make(map[string]string)

	for _, token := range tokens {
		tokenMap[token] = token
	}

	strSlice := make([]string, 0)

	if tokenMap[input] != "" {
		return true
	}

	for i := 0; i < size; i++ {
		if tokenMap[string(input[i])] != "" {
			strSlice = append(strSlice, string(input[i]))
			continue
		} else {
			oldLen := len(strSlice)
			toBeAdded := ""
			for j := i + 1; j < size; j++ {
				if tokenMap[string(input[i:j+1])] != "" {
					toBeAdded = string(input[i : j+1])
				}
			}
			if toBeAdded != "" {
				strSlice = append(strSlice, toBeAdded)
			}
			if oldLen != len(strSlice) {
				theString := strings.Join(strSlice, "")
				i = len(theString) - 1
			}

		}
	}
	fmt.Println(strSlice)
	newString := strings.Join(strSlice, "")

	if newString == input {
		return true
	}

	return false
}

// Given a non-empty array of integers, every element appears twice except for one. Find that single one.

// Note:

// Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?

// Example 1:

// Input: [2,2,1]
// Output: 1
// Example 2:

// Input: [4,1,2,1,2]
// Output: 4

func singleNumber(nums []int) int {
	frequency := make(map[int]int)

	for _, num := range nums {
		frequency[num]++
	}

	for num, times := range frequency {
		if times == 1 {
			return num
		}
	}

	return 0
}

// 1. XOR of a number with itself is 0
// 2. XOR of a number with 0 is the number itself
// 3. XOR is commutative so the order of numbers does not matter
func optimalSingleNumber(nums []int) int {
	x := 0

	for _, num := range nums {
		x ^= num
	}

	return x
}

// Write an algorithm to determine if a number is "happy".

// A happy number is a number defined by the following process: Starting with any positive integer, replace the number by the sum of the squares of its digits, and repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1. Those numbers for which this process ends in 1 are happy numbers.

// Example:

// Input: 19
// Output: true
// Explanation:
// 12 + 92 = 82
// 82 + 22 = 68
// 62 + 82 = 100
// 12 + 02 + 02 = 1

func isHappy(n int) bool {
	newNumber := 0
	visited := make(map[int]struct{})
	for {
		splitedNumber := strings.Split(strconv.Itoa(n), "")

		for _, num := range splitedNumber {
			formatedNumber, _ := strconv.Atoi(num)

			newNumber += int(math.Pow(float64(formatedNumber), 2.0))
		}

		if newNumber == 1 {
			return true
		}

		n = newNumber

		if _, ok := visited[n]; ok {
			return false
		}

		visited[newNumber] = struct{}{}

		newNumber = 0
	}
}
func main() {

	//matrix := [][]int{{4, 8}, {6, 10}}

	fmt.Println(optimalSingleNumber([]int{4, 1, 2, 1, 2}))
}
