package main

/*
● Read the following
● run the code
● read through the code
● try to understand the code

Go provides a built-in testing framework that simplifies the process of testing Go code. Here is
an overview of the file structure, naming conventions, and how testing works in Go:
File Structure and Naming Conventions:
1. Test files: Test files live in the same package as the code they test. They are named with
the following convention: `filename_test.go` where filename is the name of the file that
contains the code you want to test.
2. Test functions: Test functions must start with the word `Test` followed by a word that starts
with a capital letter. The signature of a test function is `func TestXxx(*testing.T)`, where Xxx
does not start with a lowercase letter.
*/

import "fmt"

func main() {
	fmt.Println(Add(3, 4))
}
func Add(a int, b int) int {
	return a + b
}
