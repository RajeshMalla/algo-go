package main

import (
	"fmt"
	"github.com/algo-go/binarysearch"
)

func main() {
	fmt.Println("Hello World !! ")
	elements := []int{}
	elements = append(elements, 1, 2, 3, 4, 5, 6, 77, 88, 90)

	elementIndex, err := binarysearch.BinarySearch(elements, -1)

	if err != nil {
		fmt.Println("---- Error --- ")
	}

	fmt.Println("element found at index : ", elementIndex)
}
