//formatted print
package main

import "fmt"

func main(){
	fmt.Printf("My weight on the surface of Mars is %v kg.", 50*0.3783) //%v is replaced with the result of  50*0.3783
	fmt.Printf("\nMy weight on the surface of the %v is %v kg. \n","Earth",50) // \n symbol breaks into a new line 
}


