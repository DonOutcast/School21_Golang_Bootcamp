package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	arrayNumbers := make([]int, 0)
	var median float64
	var lengthArray int
	for sc.Scan() {
		text := sc.Text()
		temp, _ := strconv.Atoi(text)
		arrayNumbers = append(arrayNumbers, temp)
	}
	sort.Ints(arrayNumbers)
	lengthArray = len(arrayNumbers)
	if lengthArray%2 != 0 {
		median = float64(arrayNumbers[lengthArray/2])
	} else {
		median = float64(arrayNumbers[lengthArray/2]+arrayNumbers[lengthArray/2-1]) / 2
	}
	fmt.Println("Median:", median)

}
