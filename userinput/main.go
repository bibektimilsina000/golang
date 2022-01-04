package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// var mystring string
	// fmt.Scanln(&mystring)
	// fmt.Println(mystring)

	// var name string = "bibek"
	// fmt.Println(name)
	// var a, _ = fmt.Println(name)
	// fmt.Println(a)

	welcome := "welcome to input output program"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("inter the rating for the movie")
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks far rating ", input)
	fmt.Printf("The type of this rating is %T", input)
}
