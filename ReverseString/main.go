package main

import "fmt"

func revreseString(str string) string {

	var reversedString string
	for i := len(str) - 1; i >= 0; i-- {
		reversedString += string(str[i])
	}
	return reversedString
}

func main() {

	str := "Sreehari"

	fmt.Println(revreseString(str))
}
