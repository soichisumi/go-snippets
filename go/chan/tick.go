package main

import (
	"time"
	"fmt"
)

func RunTick(){
	ticker:=time.NewTicker(time.Second*2)
	defer ticker.Stop()

	done := make(chan bool)
	go func(){
		time.Sleep(20*time.Second)
		done <- true
	}()
	for {
		select{
		case <- done:
			fmt.Println("done")
			return
		case t := <- ticker.C:
			fmt.Println("Current time:", t)
		}
	}
}