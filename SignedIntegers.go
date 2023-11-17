package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	//Declare a int 8
	var a int8 = 2
	var b int16 = 8
	var c int32 = 23
	var d int64 = 23
	//Size of every variables
	fmt.Printf("%d bytes\n", unsafe.Sizeof(a))
	fmt.Printf("a's type is %s\n", reflect.TypeOf(a))
	fmt.Printf("%d bytes\n", unsafe.Sizeof(b))
	fmt.Printf("a's type is %s\n", reflect.TypeOf(b))
	fmt.Printf("%d bytes\n", unsafe.Sizeof(c))
	fmt.Printf("a's type is %s\n", reflect.TypeOf(c))
	fmt.Printf("%d bytes\n", unsafe.Sizeof(d))
	fmt.Printf("a's type is %s\n", reflect.TypeOf(d))
}
