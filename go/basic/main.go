package main

import "fmt"


func manipulateByPointer(){
	f := func(v *uint32){
		*v += 5
	}

	var x uint32 = 5
	fmt.Printf("before x: %d\n", x)

	f(&x)

	fmt.Printf("after x: %d\n", x)
}

// Range で生成される変数は for内で使い回される
// このため、rangeで生成される変数をキャプチャしてはいけない
func valuesOfRangeSyntax(){
	s := []string{"a", "b", "c", "d"}
	for i, v := range s {
		fmt.Printf("loop %d\n", i)
		fmt.Printf("i: %d, v: %s\n", i, v)
		fmt.Printf("address i: %+v, v: %+v\n", &i, &v)
	}
}

// value が ゼロ値のときに加算できるんだっけという話が出たので
func maptest() {
	m1 := make(map[string]int)
	m2 := make(map[string]int)

	s := []string{ "yo", "hi", "yo", "hi"}
	for _, v := range s {
		switch v {
		case "yo":
			m1["yoresult"] += 1
		case "hi":
			m2["hiresult"] += 1
		}
	}
	fmt.Printf("m1: %+v, m2: %+v", m1, m2)
}

func main(){
	//manipulateByPointer()
	//valuesOfRangeSyntax()
	maptest()
}
