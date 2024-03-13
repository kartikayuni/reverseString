package main

import "fmt"

func main() {
	bytes := []byte{'h', 'e', 'l', 'l', 'a', 'a', 'y', 'o'}
	fmt.Println(string(bytes))
	reverseString(bytes)
	fmt.Println(string(bytes))
}
func reverseString(s []byte) {
	/*
		fmt.Println(string(s))
		s[0], s[len(s)-1-0] = s[len(s)-1-0], s[0]
		fmt.Println(string(s))
		s[1], s[len(s)-1-1] = s[len(s)-1-1], s[1]
		fmt.Println(string(s))
		s[2], s[len(s)-1-2] = s[len(s)-1-2], s[2]
		fmt.Println(string(s))
	*/
	for i := 0; i < len(s)/2; i = i + 1 {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}
