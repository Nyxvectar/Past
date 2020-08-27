package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// Mem Err.

func main() {
	scanner := bufio.NewScanner(bufio.NewReaderSize(nil, 1024))
	scanner.Scan()
	id := scanner.Text()
	valid, _ := isValidIDCard(id)
	fmt.Println(valid)
}

func isValidIDCard(id string) (bool, error) {
	if len(id) != 18 {
		return false, nil
	}
	for i := 0; i < 17; i++ {
		if !unicode.IsDigit(rune(id[i])) {
			return false, nil
		}
	}
	lastChar := rune(id[17])
	lastStr := string(lastChar)
	lastStr = strings.ToUpper(lastStr)

	if lastStr != "X" && !unicode.IsDigit(lastChar) {
		return false, nil
	}

	dateStr := id[6:14]
	year, _ := strconv.Atoi(dateStr[:4])
	if year < 1900 || year > time.Now().Year() {
		return false, nil
	}
	month, _ := strconv.Atoi(dateStr[4:6])
	if month < 1 || month > 12 {
		return false, nil
	}
	day, _ := strconv.Atoi(dateStr[6:8])
	days := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if month == 2 && (year%4 == 0 && (year%100 != 0 || year%400 == 0)) {
		days[1] = 29
	}
	if day < 1 || day > days[month-1] {
		return false, nil
	}
	weights := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	sum := 0
	for i := 0; i < 17; i++ {
		n, _ := strconv.Atoi(string(id[i]))
		sum += n * weights[i]
	}
	checkCodes := []string{"1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2"}
	return checkCodes[sum%11] == lastStr, nil
}
