// true or false values

package main
import (
	"fmt"
	"strings"
	)
func main() {
	fmt.Println("You find yourself in a dimly lit cavern.")
	var command = "walk outside" // try deleting the word outside see what happens
	var exit = strings.Contains(command, "outside")
	fmt.Println("You leave the cave:", exit)
}
