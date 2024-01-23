package main

import (
	"fmt"

	"github.com/alexandr-orlov/corecatcher/pkg/greetings"
)

func main() {
	fmt.Println("Hi go")
	mess := greetings.Hello("dddddddd")
	fmt.Println(mess)
}
