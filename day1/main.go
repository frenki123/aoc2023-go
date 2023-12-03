package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"
)

var line = 0

func updateSum(num string, sum int) (int, error) {
	line += 1
	if len(num) < 2 {
		num += num
	}
	num_val, err := strconv.Atoi(num)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Line:", line)
	fmt.Println("Current num:", num_val)
	sum += num_val
	return sum, nil
}

func main() {
	file_bytes, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	sum := 0
	found_first_num := false
	num_letter := ""
	for _, c := range file_bytes {
		if string(c) == "\n" {
			sum, _ = updateSum(num_letter, sum)
			found_first_num = false
			num_letter = ""
		} else if unicode.IsDigit(rune(c)) && !found_first_num {
			num_letter += string(c)
			found_first_num = true
		} else if unicode.IsDigit(rune(c)) && found_first_num {
			if len(num_letter) >= 2 {
				num_letter = num_letter[:1] + string(c)
			} else {
				num_letter += string(c)
			}
		}
	}
	sum, _ = updateSum(num_letter, sum)
	fmt.Println("Sum is:", sum)
}
