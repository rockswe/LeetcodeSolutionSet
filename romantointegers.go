package main

func romanToInt(s string) int {
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	
	result := 0
	for i := (len(s) - 1); i >= 0; i-- {
		if (i == (len(s) - 1)) || (romanMap[s[i]] >= romanMap[s[i+1]]) {
			result += romanMap[s[i]]
		} else {
			result -= romanMap[s[i]]
		}
	}

	return result
}
