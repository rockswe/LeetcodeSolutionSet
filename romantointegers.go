package main

import "fmt"

func romanToInt(s string) int {

	romanCharacters := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	total := 0

	for i := 0; i < len(s); i++ {
		if i < len(s)-1 && romanCharacters[string(s[i])] < romanCharacters[string(s[i+1])] {
			total -= romanCharacters[string(s[i])]
		} else {
			total += romanCharacters[string(s[i])]
		}
	}

	return total
}

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("IV"))
	fmt.Println(romanToInt("IX"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}
