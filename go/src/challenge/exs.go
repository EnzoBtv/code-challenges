package main

/*
 * Auxiliar functions
 */

func pushIndex(numberArr *[]int, index, number int) {
	*numberArr = append(*numberArr, 0)

	copy((*numberArr)[index+1:], (*numberArr)[index:])

	(*numberArr)[index] = number
}

func findInArray(tokens []string, input string) bool {
	for _, token := range tokens {
		if token == input {
			return true
		}
	}
	return false
}

/*
 * TASK 1
 * Using Gocyclo tool it defines the cyclomatic complexity as 5
 */

func compareStringWithSlice(input string, wordSlice []string) bool {
	size := len(input)
	boolSlice := make([]bool, size+1)
	boolSlice[0] = true
	for i := 1; i <= size; i++ {
		for j := 0; j < i; j++ {
			boolSlice[i] = boolSlice[j] && findInArray(wordSlice, input[j:i])
			if boolSlice[i] {
				break
			}
		}
	}
	return boolSlice[size]
}

/*
 * TASK 2
 * Using Gocyclo tool it defines the cyclomatic complexity as 6
 */
func mergeIntervalLists(intervalList [][]int) []int {
	numberFrequency := make(map[int]int)
	for i, interval := range intervalList {
		lastNumber := interval[len(interval)-1]
		for j := 0; j < len(interval); j++ {
			number := interval[j]
			if number+1 == lastNumber {
				numberFrequency[number]++
				numberFrequency[lastNumber]++
				break
			}
			numberFrequency[number]++

			pushIndex(&interval, j+1, number+1)

		}

		intervalList[i] = interval
	}

	duplicateNumbers := make([]int, 0)

	for key, value := range numberFrequency {
		if value > 1 {
			duplicateNumbers = append(duplicateNumbers, key)
		}
	}

	return duplicateNumbers
}
