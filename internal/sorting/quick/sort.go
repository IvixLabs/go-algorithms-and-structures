package quick

import (
	"cmp"
	"math/rand"
)

func Sort[T cmp.Ordered](items []T) []T {

	ln := len(items)
	if ln < 2 {
		return items
	}

	pivotIdx := rand.Intn(ln)
	pivot := items[pivotIdx]

	maxArrPartLen := ln - 1
	less := make([]T, 0, maxArrPartLen)
	greater := make([]T, 0, maxArrPartLen)

	for i, v := range items {
		if i == pivotIdx {
			continue
		}

		if v < pivot {
			less = append(less, v)
		}

		if v > pivot {
			greater = append(greater, v)
		}
	}

	res := make([]T, 0, ln)
	res = append(res, Sort(less)...)
	res = append(res, pivot)
	res = append(res, Sort(greater)...)

	return res
}
