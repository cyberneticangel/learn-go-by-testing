/*

Simple Greeter

*/

package hello

import (
	"fmt"
)

func greeting(user string) string {
	return "Hello " + user
}

func main() {
	fmt.Println("Hello, go!")
	fmt.Println("Hello World")
	fmt.Println(greeting("Dylan")) // this will fail
	fmt.Println(greeting("World")) // this will pass

}
