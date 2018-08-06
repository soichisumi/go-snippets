package golang

import (
	"time"
	"fmt"
)

func say(s string){
	for i:=0;i<5;i++{
		time.Sleep(100*time.Millisecond)
		fmt.Println(s)
	}
}

func RunGor(){
	go say("goroutine")
	say("for loop")
}
