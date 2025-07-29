package main

import "fmt"

func push(arr []int, val int) []int {

	arr = append(arr, val)
	return arr
}

func pop(arr []int) []int {

	arr = arr[:len(arr)]
	return arr
}

func peek(arr []int) int {
	val := arr[len(arr)-1]
	return val
}

func main() {

	arr := []int{10, 20, 30}
	fmt.Println(push(arr, 40))

	fmt.Println(pop(arr))

	fmt.Println(peek(arr))
}
