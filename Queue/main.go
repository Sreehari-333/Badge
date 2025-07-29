package main

import "fmt"

func enQueue(arr []int, val int) []int {
	arr = append(arr, val)
	return arr
}

func deQueue(arr []int) []int {
	arr = arr[1:]
	return arr
}

func peek(arr []int) int {
	val := arr[0]
	return val
}

func main() {
	arr := []int{1, 2, 3, 4, 5}

	fmt.Println(enQueue(arr, 6))
	fmt.Println(deQueue(arr))
	fmt.Println(peek(arr))
}
