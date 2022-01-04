package main

import "fmt"

var (
	spiderman = "i am spider man"
	age       = 18
	powers    = "spider sence, webshooters"
)

func main() {

	var i int
	i = 42
	i = 78
	//next way to declaro variable

	var j int = 21

	//aother way
	k := 25

	//strings

	var batman string = "i am bat man"

	var superman string
	superman = "i am superman"

	thor := "i am thor "

	var antman, hulk string = "hi i am antman ", "i am hulk"

	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)

	//printing strings

	fmt.Println(superman, age, powers)
	fmt.Println(batman)
	fmt.Println(superman)
	fmt.Println(thor)
	fmt.Println(hulk)
	fmt.Println(antman)
}
