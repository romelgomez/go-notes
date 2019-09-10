package main

import (
	"fmt"
)

func main() {
	askNane()
}

func askNane() {
	var name string
	fmt.Println("What is your name?")
	fmt.Scanf("%s", &name)
	fmt.Printf("Fuck you %s! \n", name)
}
