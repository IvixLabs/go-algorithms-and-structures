package bubble

import "cmp"

func Sort[T cmp.Ordered](items []T) []T {
	n := len(items)
	sortedItems := make([]T, n)
	copy(sortedItems, items)

	for i := 0; i < n; i++ {
		swapped := false

		for j := 0; j < n-i-1; j++ {
			if sortedItems[j] > sortedItems[j+1] {
				t := sortedItems[j]
				sortedItems[j] = sortedItems[j+1]
				sortedItems[j+1] = t
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}

	return sortedItems
}
