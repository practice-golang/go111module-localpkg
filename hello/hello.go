package main // import "hello"

import (
	"fmt"

	"mypkg" // A local package outside of main package

	"spkg" // Sub-package
)

func main() {
	// Hello world
	fmt.Println("Hello world!")

	var myName = "Chulsoo"
	yourName := "Younghee"

	// Using package
	fmt.Println(mypkg.SayHello())
	var name string
	name = mypkg.MyName(myName)
	fmt.Println(name)

	fmt.Println(mypkg.YourName(yourName))

	fmt.Println(spkg.SayGoodToSeeYou())
}
