// Using the concept of Array and looping in GoLang:
// create an array checking program to
// compare the contents of 2 arrays and explain each step! (LO3, 30%)

// 1. input length of array v
// 2. make looping input for 2 arrays
// 3. make looping for checking

package main

import "fmt"

func main() {
	var length int
	fmt.Print("Input the length of array: ")
	fmt.Scan(&length)

	var arr1 []string
	var arr2 []string
	var n string

	fmt.Println("Input array 1:")
	for i := 0; i < length; i++ {
		fmt.Scan(&n)
		arr1 = append(arr1, n)
	}

	fmt.Println("Input array 2:")
	for i := 0; i < length; i++ {
		fmt.Scan(&n)
		arr2 = append(arr2, n)
	}

	for i := 0; i < length; i++ {
		if arr1[i] != arr2[i] {
			fmt.Printf("Array ke %d berbeda\n", i+1)
		}
	}

}
