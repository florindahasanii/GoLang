// if else conditionals

package main
import "fmt"

func main(){
	var command = "go east"
	//var command = "go inside"
	if command == "go east"{
		fmt.Println("You head further up the mountain.")
		} else if command == "go inside"{
		fmt.Println("You enter the cave where you live out the rest of your life.")
		}else{
		fmt.Println("You are lost.")
		}
	}
