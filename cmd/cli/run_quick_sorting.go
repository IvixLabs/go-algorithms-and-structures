package main

import (
	"ivixlabs.com/go-algorithms-and-structures/internal/sorting/quick"
	"log"
	"math/rand/v2"
)

func runQuickSorting() {
	arrLen := rand.IntN(100)

	arr := make([]int, arrLen)
	for i := range arrLen {
		arr[i] = rand.IntN(100)
	}

	log.Println("Income array:", arr)

	sortArr := quick.Sort(arr)

	log.Println("Sorted array:", sortArr)
}
