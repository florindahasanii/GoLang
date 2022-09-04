// a countdown where every second there’s a 1 in 100 chance that the launch fails

package main
import( 
	"fmt" 
	"time"
	"math/rand"
	)

func main(){
	var count = 10 
	for count > 0 {
		fmt.Println(count)
		time.Sleep(time.Second)
		if rand.Intn(100) == 0 {
			break 
			}
			count--
		}
		if count == 0 {
			fmt.Println("Lift Offff!")
		}else { 
			fmt.Println("Launch failed.")
		}
}
