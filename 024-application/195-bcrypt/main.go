package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := "password123"
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(bs)

	login := "password1234"

	err = bcrypt.CompareHashAndPassword(bs, []byte(login))
	if err != nil {
		fmt.Println("Wrong password")
		return
	}
	fmt.Println("You are logged in")
}
