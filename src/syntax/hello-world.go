package main //GO programs are organized as packages
import (
	"fmt"
	"runtime"
)

// The import statement allows you to use external code.
// The fmt package provided by the standard library allows you
// to format and output data

func main() {
	fmt.Println("Hello world!!!")
	fmt.Println(runtime.NumCPU())
}
// The main function is what gets executed when you run your application - just like in C


