// declaring multiple variables in a group or each one on its own line

package main

import "fmt"

func main(){
	var a = 10
	var b = 5
	
	var (
		c = 10
		d = 5
	    )
	fmt.Println("The rectangle has an area of: ", a*b)
	
	fmt.Println("The rectangle has a perimeter of: ", a+b+c+d)
	
	}
