package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)
func main(){
	tmp := 0
	inputFile, err := os.Open("1.txt")
	scanner := bufio.NewScanner(inputFile)
	if(err != nil){
		fmt.Print(err)
	}
	defer inputFile.Close()
	for scanner.Scan(){
		str := scanner.Text()
		num, _ := strconv.Atoi(str)
		tmp = tmp+num
	}
	fmt.Print(tmp)
	if err := scanner.Err(); err != nil {
		fmt.Print(err)
	}
}