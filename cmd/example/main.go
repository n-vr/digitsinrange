package main

import (
	"fmt"

	"github.com/n-vr/digitsinrange"
)

func main() {
	fmt.Println("5s in 0 to 10:", digitsinrange.CountDigitOccurancesInRangeNaive(5, 10))
	fmt.Println("5s in 0 to 20:", digitsinrange.CountDigitOccurancesInRangeNaive(5, 20))
	fmt.Println("5s in 0 to 30:", digitsinrange.CountDigitOccurancesInRangeNaive(5, 30))
	fmt.Println("5s in 0 to 40:", digitsinrange.CountDigitOccurancesInRangeNaive(5, 40))
	fmt.Println("5s in 0 to 50:", digitsinrange.CountDigitOccurancesInRangeNaive(5, 50))
	fmt.Println("5s in 0 to 60:", digitsinrange.CountDigitOccurancesInRangeNaive(5, 60))
	fmt.Println("5s in 0 to 70:", digitsinrange.CountDigitOccurancesInRangeNaive(5, 70))
	fmt.Println("5s in 0 to 80:", digitsinrange.CountDigitOccurancesInRangeNaive(5, 80))
	fmt.Println("5s in 0 to 90:", digitsinrange.CountDigitOccurancesInRangeNaive(5, 90))
	fmt.Println("5s in 0 to 100:", digitsinrange.CountDigitOccurancesInRangeNaive(5, 100))
	fmt.Println("5s in 0 to 110:", digitsinrange.CountDigitOccurancesInRangeNaive(5, 110))

	fmt.Println()

	fmt.Println("5s in 0 to 10:", digitsinrange.CountDigitOccurancesInRange(5, 10))
	fmt.Println("5s in 0 to 20:", digitsinrange.CountDigitOccurancesInRange(5, 20))
	fmt.Println("5s in 0 to 30:", digitsinrange.CountDigitOccurancesInRange(5, 30))
	fmt.Println("5s in 0 to 40:", digitsinrange.CountDigitOccurancesInRange(5, 40))
	fmt.Println("5s in 0 to 50:", digitsinrange.CountDigitOccurancesInRange(5, 50))
	fmt.Println("5s in 0 to 60:", digitsinrange.CountDigitOccurancesInRange(5, 60))
	fmt.Println("5s in 0 to 70:", digitsinrange.CountDigitOccurancesInRange(5, 70))
	fmt.Println("5s in 0 to 80:", digitsinrange.CountDigitOccurancesInRange(5, 80))
	fmt.Println("5s in 0 to 90:", digitsinrange.CountDigitOccurancesInRange(5, 90))
	fmt.Println("5s in 0 to 100:", digitsinrange.CountDigitOccurancesInRange(5, 100))
	fmt.Println("5s in 0 to 110:", digitsinrange.CountDigitOccurancesInRange(5, 110))

	fmt.Println("5s in 0 to 38469:", digitsinrange.CountDigitOccurancesInRange(5, 38469))

}
