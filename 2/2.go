package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)
func main(){
	tmp := 0
	sign := ""
	var value []string
	sol := 0
	var numArray []int

	inputFile, err := os.Open("2.txt")
	scanner := bufio.NewScanner(inputFile)
	if(err != nil){
		fmt.Print(err)
	}
	defer inputFile.Close()
	for scanner.Scan(){
		str := scanner.Text()
		sign := str[0]
		value := str[1:]
		num, _ := strconv.Atoi(value)
		if sign == '-' {
			tmp = tmp - num
			numArray = append(numArray, tmp)
		} else {
			tmp = tmp + num
			numArray = append(numArray, tmp)
		}
	}
	for i := 0; i < len(numArray); i++ {
		for j := i + 1; j < len(numArray); j++ {
			if(numArray[i] == numArray[j]){
				sol = numArray[i]
				break
			}
		}
	}
	fmt.Print(sol)
	if err := scanner.Err(); err != nil {
		fmt.Print(err)
	}
}