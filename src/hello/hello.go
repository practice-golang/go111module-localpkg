package main

import (
	"fmt"

	"mypkg" // 메인 패키지 바깥에 존재하는 로컬 패키지

	"hello/spkg" // Sub-package
	"vpkg"       // Vendor Package : must be imported as vpkg not /hello/vendor/vpkg
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

	fmt.Println(spkg.SayGoodToSeeYou()) // 이제는 된다.

	fmt.Println(vpkg.SayGoodbye()) // 이제는 안 된다.
}
