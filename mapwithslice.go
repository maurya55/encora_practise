// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main

import (
	"fmt"
)

func main() {
	//   x := map[string]int{"one":1,"two":2,"three":3}
	var y = make([]map[string]int, 0)
	x := make(map[string]int)
	x["one"] = 23
	x["two"] = 232
	y = append(y, x)
	fmt.Print(x)
	// for index,val:= range x{
	//     fmt.Println(index,val)
	// }

}
