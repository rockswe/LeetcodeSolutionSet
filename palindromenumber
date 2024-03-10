package main

import "fmt"

func isPalindrome(x int) bool {
    if x < 0 || (x != 0 && x%10 == 0) {
        return false
    }
    revertedNumber := 0
    for x > revertedNumber {
        revertedNumber = revertedNumber*10 + x%10
        x /= 10
    }
    return x == revertedNumber || x == revertedNumber/10
}

func main() {
    fmt.Println(isPalindrome(121)) // true
    fmt.Println(isPalindrome(-121)) // false
    fmt.Println(isPalindrome(10)) // false
}
