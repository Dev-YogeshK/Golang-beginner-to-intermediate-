package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	//Declare a float64
	var a float32 = 4
	var b float64 = 2

	//Size in bytes
	fmt.Printf("%d bytes\n", unsafe.Sizeof(a))
	fmt.Printf("a's type is %s\n", reflect.TypeOf(a))

	fmt.Printf("%d bytes\n", unsafe.Sizeof(b))
	fmt.Printf("a's type is %s\n", reflect.TypeOf(b))
	
	//Default is float64 when you don't specify a type
	c := 3.2
	fmt.Printf("b's type is %s\n", reflect.TypeOf(c))
}
