package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func inputNumber(array *[]int) {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		text := sc.Text()
		temp, _ := strconv.Atoi(text)
		*array = append(*array, temp)
	}
}

func main() {
	arrayNumbers := make([]int, 0)
	var median float64
	var lengthArray int
	inputNumber(&arrayNumbers)
	sort.Ints(arrayNumbers)
	fmt.Println(arrayNumbers)
	lengthArray = len(arrayNumbers)
	if lengthArray%2 != 0 {
		median = float64(arrayNumbers[lengthArray/2])
	} else {
		median = float64(arrayNumbers[lengthArray/2]+arrayNumbers[lengthArray/2-1]) / 2
	}
	fmt.Println("Median:", median)

}
