// generating random numbers between 1 and 10
package main

import (
	"fmt"
	"math/rand"
	)
func main() {
	var num = rand.Intn(10) + 1 // remeber to add 1 because if you forget it will give you a number between 0-9 off-by-one error
	fmt.Println(num) 
	num = rand.Intn(5) + 1 // the Intn function is prefixed with the package name rand
	fmt.Println(num)
}
