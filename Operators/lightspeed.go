//How long does it take to get to Mars? 

package main
import "fmt"

func main(){
	const lightSpeed = 299792 // km/s
	var distance = 56000000 // km
	fmt.Println(distance/lightSpeed,"seconds")
	distance = 401000000
	fmt.Println(distance/lightSpeed, "seconds")
  }
//The first calculation is based on Mars and Earth being nearby, with distance declared and
//assigned an initial value of 56,000,000 km. Then the distance variable is assigned a new
//value of 401,000,000 km, with the planets on opposite sides of the Sun, though plotting a
//course directly through the Sun could be problematic.
