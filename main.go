package main // import "hello"

import (
	"hello/mypkg"
	"log"
)

func main() {
	log.Println(mypkg.MyName("철수"))
	log.Println(mypkg.YourName("영희"))
	log.Println(mypkg.SayHello())
}
