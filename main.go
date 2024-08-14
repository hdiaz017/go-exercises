package main

import (
	"fmt"

	"github.com/GoesToEleven/puppy"
	"github.com/hdiaz17/sandbox/main/hello"
)

var number int

const name string = "luis"

func main() {
	s1 := puppy.Bark()
	s2 := puppy.Barks()
	puppy.From12()

	fmt.Println(s1, s2)
	fmt.Println(number, name)
	hello.Hello()
}
