package binarysearch

import "fmt"

func BinarySearch(elements []int, element int) (index int, err error) {
	input := fmt.Sprintf("Elements Array %v, Element index to found %d ", elements, element)

	fmt.Println(input)

	left := 0
	right := len(elements) - 1

	/*
	   If the element is there on array, somehow it will reach stage where element == array[mid]
	*/
	for left <= right {
		mid := left + (right-left)/2
		if elements[mid] == element {
			return mid, nil
		}
		if elements[mid] < element {
			left = mid + 1

		} else if elements[mid] > element {
			right = mid - 1
		}
	}

	return -1, nil
}
