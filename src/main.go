package main

import "fmt"

func main() {
	//lexer
	fileName := "test.txt"
	Pretreatment(fileName)
	recognition("test.temp")
	for k, v := range signlist {
		fmt.Println("(", k, ",", v, ")")
	}
}


