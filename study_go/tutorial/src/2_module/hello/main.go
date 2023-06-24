package main

import (
	"fmt"
	"tutorial/2_module/greetings"
)

func main() {
	message := greetings.Hello("Dowon")
	fmt.Println(message)
}
