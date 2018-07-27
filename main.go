package main

import (
	"fmt"
	"reflect"
	"github.com/yoyoyousei/GoSandBox/communication"
)

//type test struct {
//	a string
//	b float64
//}

func testAssert(arg interface{}){
	switch v:=arg.(type) {
	case string:
		fmt.Println("string")
	case int, int32, int64, uint32, uint64:
		fmt.Println("ints / type is : ", reflect.TypeOf(v))
	default:
		fmt.Println("default / type is :", reflect.TypeOf(v))
	}
	fmt.Println("target: ", arg)
	x, ok :=arg.(int)
	fmt.Printf("x: %v, ok = %v", x, ok)
	y := arg.(uint64)
	fmt.Printf("test cast %s: ", reflect.TypeOf(y))
}

func interf(x interface{}){
	fmt.Printf("interface : %s\n", x)
	xs, _ := x.(string)
	fmt.Printf("string : %s\n", xs)
	//xf, _ := x.(float64)
	fmt.Printf("conv to str, %s", fmt.Sprint(x))
}

func main() {
	//val := test{a:"yoyo", b:3.14}
	//fmt.Println(val.a)
	//concurrency.RunGor()
	//concurrency.RunChan2()
	//ext.RunTick()
	//communication.RunGori()
	//communication.RunRPC()
	//var x interface{}
	//x = 5
	//switch v:=x.(type) {
	//case string:
	//	fmt.Println("string")
	//case int, int32, int64, uint32, uint64:
	//	fmt.Println("ints / type is : ", reflect.TypeOf(v))
	//default:
	//	fmt.Println("default / type is :", reflect.TypeOf(v))
	//}
	//fmt.Println("target: ", x)
	//var x interface{}
	//x = 40000000000
	//_, ok := x.(int)
	//fmt.Println("res; ", ok)
	//
	//testAssert(x)
	//x:=555
	//fmt.Printf("x: type: %s , size: %d", reflect.TypeOf(x), reflect.TypeOf(x).Size())

	//x := 1.4
	//x2 := int(x)
	//fmt.Printf("%f, %d\n", x, x2)
	//interf(x)
	//x := 5.235
	//
	//fmt.Printf("%d, %d, %d, %f, %f, %f, %s, %s, %s", int(4000000000), float64(5000000000), float64(5.234), int(400000000), float64(4.23424), float64(0.00052354234), int(4000000000), int(x), float64(5.234))
	communication.RunGori()
}

