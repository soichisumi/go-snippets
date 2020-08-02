package main

import (
	"fmt"
	"sort"
	"time"
)

func channelBasic() {
	// closeFirst
	c := make(chan int, 4)
	for _, v := range []int{ 1, 4, 25, 52 } {
		c <- v
	}
	close(c)
	for v := range c {
		fmt.Printf("%d received\n", v)
	}
}

func sortSlice(){
	type Dog struct {
		Name    string
		Age     int
		BoredAt time.Time
	}
	now := time.Now()
	t1 := now
	t2 := now.Add( -2 * time.Hour)
	t3 := now.Add( 1 * time.Hour)
	dogs := []Dog{
		Dog{Name: "yo1", Age: 1, BoredAt: t1},
		Dog{Name: "bow", Age: 1, BoredAt: t2},
		Dog{Name: "wow", Age: 1, BoredAt: t3},
	}
	fmt.Printf("pre: %+v\n", dogs)
	sort.Slice(dogs, func(i, j int) bool {
		return dogs[i].BoredAt.Unix() < dogs[j].BoredAt.Unix()
	})
	fmt.Printf("post: %+v\n", dogs)
}

func main(){
	//channelBasic()
	sortSlice()
}