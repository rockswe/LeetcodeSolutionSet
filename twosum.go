package main

import "fmt"

func twoSum(nums []int, target int) []int {
    m := make(map[int]int) // Create a map to store value to index mapping
    for i, num := range nums {
        complement := target - num
        if j, ok := m[complement]; ok {
            // If the complement is found, return the indices
            return []int{j, i}
        }
        // Store the index of the current number in the map
        m[num] = i
    }
    // Return an empty slice if no solution is found
    return []int{}
}

func main() {
    nums := []int{2, 7, 11, 15}
    target := 9
    result := twoSum(nums, target)
    fmt.Println(result) // Output: [0 1], because nums[0] + nums[1] = 2 + 7 = 9
}
