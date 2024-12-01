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
	var times []int
	l := 0
	r := 0
	count := 0
	var store map[int]int
	store = make(map[int]int)
	for l < len(leftnums){
		if value , ok := store[leftnums[l]]; ok {
			times = append(times,value * leftnums[l])
			l++
			
		}else if leftnums[l] < rightnums[r]{
			store[leftnums[l]] = count
			times = append(times,count * leftnums[l])
			l++
			count = 0
		}else if leftnums[l] == rightnums[r] {
			count += 1
			r++
		}else {
			r++

		}
	}

	similarityScore := 0
	for _, nums := range times {
		similarityScore += nums
	}
	fmt.Println(similarityScore)


	readFile.Close()
}
