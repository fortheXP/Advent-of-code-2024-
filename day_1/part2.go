package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	readFile , err := os.Open("input.txt")
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
