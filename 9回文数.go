package main

import "strconv"
import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	println(s[1])
	fmt.Printf("a,%T", s)
	return true
}

func main() {
	print(isPalindrome(121121212))
}
