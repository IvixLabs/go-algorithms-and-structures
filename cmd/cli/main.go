package main

import (
	"flag"
)

var (
	quickSorting  *bool
	bubbleSorting *bool
)

func init() {
	quickSorting = flag.Bool("quick-sorting", false, "run quick sorting")
	bubbleSorting = flag.Bool("bubble-sorting", false, "run bubble sorting")
}

func main() {
	flag.Parse()

	switch true {
	case *quickSorting:
		runQuickSorting()
	case *bubbleSorting:
		runBubbleSorting()
	default:
		flag.PrintDefaults()
	}

}
