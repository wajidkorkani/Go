package main

import "fmt"

var x = 0
var i = 20
var yeet = true
var noYeet = false

func functionOne() {
	for x < 20 {
		fmt.Println("x = ", x)
		if i == 20 {
			fmt.Println("i = ", i)
			i++
		} else {
			fmt.Println("i = ", i)
			i--
		}
		if yeet || noYeet {
			fmt.Println("Yeet mainia")
			yeet = !yeet
		} else {
			fmt.Println("NoYeet Mainia")
			yeet = !yeet
		}
		x++
	}
}

//This function is to take information from user
func userInfo() {
	var name string
	var age int
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)
	fmt.Print("Enter your age: ")
	fmt.Scan(&age)
	fmt.Println("Your name is", name, "Your age is", age)
}
func main() {
	userInfo()
}
