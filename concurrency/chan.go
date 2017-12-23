package concurrency

import "fmt"

func sum(s []int, c chan int) { //c chan int : intを送受信するチャンネル
	sum := 0
	for _, v := range s { //for index, ループ対象
		sum += v
	}
	c <- sum
}

func RunChan(){
	s:=[]int{7,2,8,-9,4,0}
	c:=make(chan int)
	go sum(s[:len(s)/2],c)
	go sum(s[len(s)/2:],c)
	x, y := <-c, <-c
	fmt.Println(x,y,x+y)
}