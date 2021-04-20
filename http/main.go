package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	if len(os.Args) > 1 {
		strFileName := strings.Join(os.Args[1:], " ")
	  file, error := os.Open(strFileName)
	  if error != nil {
		  fmt.Println("Error:", error)
		  os.Exit(1)
	  }
	  io.Copy(os.Stdout, file)
	}
}



/*
package main

import (
	"fmt"
	//"io"
	//"net/http"
//	"os"
)
type shape interface {
	getArea() float64 
}
type triangle struct{
	height float64
	base float64
}

type square struct{
	sideLength float64
}

type logWriter struct{}

func main() {
	/*
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}


	bs := make([]byte, 99999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))
	

	io.Copy(os.Stdout, resp.Body)

	lw := logWriter{}
	io.Copy(lw, resp.Body)

	shapeSquare := square{sideLength: 10}
	shapeTriangle := triangle{base: 10, height:10}

	printArea(shapeSquare)
	printArea(shapeTriangle)
	
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.height * t.base
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (logWriter) Write(bs []byte) (int, error){
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
*/