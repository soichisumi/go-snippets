package main

import (
	"github.com/yoyoyousei/SandBox/concurrency"
)

type test struct {
	a string
	b float64
}

func main() {
	//val := test{a:"yoyo", b:3.14}
	//fmt.Println(val.a)

	//concurrency.RunGor()

	concurrency.RunChan()

}
