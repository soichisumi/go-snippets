package main

import (
	"fmt"
	"math"
)

func sum(s []int, c chan int) { //c chan int : intを送受信するチャンネル
	sum := 0
	for _, v := range s { //for index, ループ対象
		sum += v
	}
	c <- sum
}

func RunChan() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}

func RunChan2() {
	s := []int{7, 4, 6, 21, 634, 63, 4356, 234, 1241, 4}
	cint := make(chan int, len(s))
	for _, x := range s {
		fmt.Printf("push: %d\n", x)
		cint <- x
	}
	close(cint)
	cflo := make(chan float64, len(s))
	go func(ci chan int, cf chan float64) {
		for x:= range ci{
			fmt.Printf("pop: %d\n", x)
			cf <- math.Sqrt(float64(x))
		}
		close(cf)
	}(cint, cflo)

	res := make([]float64, 0, len(s))
	for nf := range cflo {
		fmt.Printf("pop2: %f\n", nf)
		res = append(res, nf)
	}
	fmt.Printf("input: %s\n", s)
	fmt.Printf("res: %s\n", res)
}
