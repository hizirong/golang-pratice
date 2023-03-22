// code resourse: https://hackmd.io/@BKLiang/golang_ds
package main

import "fmt"

func main() {
	// Array & Slice
	arr := [5]int{1, 2, 3, 4, 5}

	for index, value := range arr {
		fmt.Printf("[%d]: %d\n", index, value)
	}

	// Map
	users := make(map[string]string)

	users["7458"] = "Charlie"
	users["6422"] = "POKA"
	users["9738"] = "Wendy"

	for key, val := range users {
		fmt.Printf("[%s]: %s\n", key, val)
	}

}
