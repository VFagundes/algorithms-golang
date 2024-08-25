package main

import "fmt"

func main() {

	fmt.Println(romanToInt("MCMXCIV"))
	fmt.Println(romanToInt("XIX"))
}
func romanToInt(s string) int {
	nums := getNumbersMap()
	i, n := 0, len(s)
	total := 0
	for i < n {
		if i < n-1 && nums[s[i+1]] > nums[s[i]] {
			total += nums[s[i+1]] - nums[s[i]]
			i += 2

			continue
		}
		total += nums[s[i]]
		i++

	}
	return total
}
func getNumbersMap() map[byte]int {
	return map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

}
