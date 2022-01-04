package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("hellou user welcome to add number program")
	fmt.Println("input a number i will add 2 in that number")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	intiger, _ := strconv.ParseFloat(strings.TrimSpace(input), 64)
	output := intiger + 2
	err := 0
	if err != 0 {
		fmt.Println(err)
	} else {
		fmt.Println("the number you input is ", intiger)
		fmt.Println("The final number after sum is ", output)
	}

}
