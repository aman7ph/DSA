package main

import (
	"fmt"

	"github.com/aman7ph/DSA/sort/insertion"
)

func main() {
	arr := []int{5, 2, 4, 6, 1, 3}

    insertion.InsertionSort(arr)
	fmt.Println(arr)

}