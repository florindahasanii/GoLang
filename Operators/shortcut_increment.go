//using shortcuts and increment operators

package main

import "fmt"

func main(){
	var years = 22
	fmt.Println("My years: " , years )
	years *=2
	fmt.Println("My fathers years: " , years) // double the years / 44
	years+=1
	fmt.Println("My fathers birthday, opss one year older: " , years) // happy birthdayy!!! / 45
	years++  
	fmt.Println("Next year he will be: " , years)  // should display 46
	years-- 
	fmt.Println("At least he's still: " , years) // now it 45 again
	}
// go does not support the prefix increment ++years
