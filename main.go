package main

import (
	"fmt"
	"main/model"
	"main/tronics"
	// "main/model"
)

func main() {
	model.StartConn()
	// fmt.Println(rand.Intn(20))
	tronics.Start()
	fmt.Println("hello there newly made ")
}
