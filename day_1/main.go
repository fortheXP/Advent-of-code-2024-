package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)
func parse(s string) int {
	num , err := strconv.Atoi(s)
	if err!= nil {
		fmt.Println(err)
	}
	return num
}
func absdifference(a , b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
func main() {
	readFile , err := os.Open("../input.txt")
	if err!= nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var leftnums []int
	var rightnums []int
	for fileScanner.Scan(){
		line := fileScanner.Text()
		split := strings.Split(line,"   ")
		lnum := parse(split[0])
		rnum := parse(split[1])
		leftnums = append(leftnums,lnum)
		rightnums = append(rightnums,rnum)
	}
	sort.Ints(leftnums)
	sort.Ints(rightnums)
	totalDistance := 0

	for i :=0;i < len(leftnums); i++{
		totalDistance += absdifference(leftnums[i],rightnums[i])

	}
	fmt.Println(totalDistance)
	rightmap := make(map[int]int)
	for _, r := range rightnums {
		rightmap[r]++
	}
	similarityScore := 0
	for _ ,l := range leftnums {
		count := rightmap[l]
		similarityScore += l * count
	}
	fmt.Println(similarityScore)


	readFile.Close()
}
