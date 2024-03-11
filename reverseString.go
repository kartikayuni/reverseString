package main

import "fmt"

func main() {
	bytes := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(bytes)
}
func reverseString(s []byte) {
	fmt.Println(string(s))
	s[0], s[len(s)-1] = s[len(s)-1], s[0]

	fmt.Println(string(s))
}
