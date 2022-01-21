package main

import "fmt"

func binary_search(list []int, itemToFind int) (position *int) {
	lowest := 0
	highest := len(list) - 1

	for lowest <= highest {
		middle := (lowest + highest) / 2
		guess := list[middle]
		if guess == itemToFind {
			position = &middle
			return position
		} else if guess > itemToFind {
			highest = middle - 1
		} else {
			lowest = middle + 1
		}
	}

	return nil
}

func main() {
	myList := []int{1, 2, 3, 4, 5, 6, 7, 8}
	pos := binary_search(myList, 4) // should return position 3 od the list
	if pos == nil {
		fmt.Println("element is not in list")
	}

	fmt.Printf("Found element in position %d of list", *pos)
}
