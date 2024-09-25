package main

import (
	"ivixlabs.com/go-algorithms-and-structures/internal/sorting/bubble"
	"log"
	"math/rand/v2"
)

func runBubbleSorting() {
	arrLen := rand.IntN(100)

	arr := make([]int, arrLen)
	for i := range arrLen {
		arr[i] = rand.IntN(100)
	}

	log.Println("Income array:", arr)

	sortArr := bubble.Sort(arr)

	log.Println("Sorted array:", sortArr)
}
