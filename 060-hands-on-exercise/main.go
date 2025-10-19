package main

import (
	"fmt"
	"github.com/GoesToEleven/dog"
	"github.com/GoesToEleven/puppy"
)

func main() {
	puppy.From13()
	z := dog.WhenGrownUp(puppy.Barks())

	fmt.Println(z)
}
