package main

import (
	"fmt"
	"github.com/emirpasic/gods/sets/treeset"
	"github.com/emirpasic/gods/utils"
	"github.com/google/uuid"
	"log"
)

func TestGods(){
	// should implement i
	type Dog struct {
		ID string
		Name string
		Age int
	}

	testdata := []Dog{
		{ID: uuid.New().String(), Name: "doggo", Age: 1},
		{ID: uuid.New().String(), Name: "catto", Age: 29345},
	}
	dogComparator := func(a, b interface{}) int {
		dogA, ok := a.(Dog)
		if !ok {
			log.Fatalf("a is not Dog. a : %+v", a)
		}
		dogB, ok := b.(Dog)
		if !ok {
			log.Fatalf("b is not Dog. b : %+v", b)
		}
		if dogA.Age == dogB.Age {
			return utils.StringComparator(dogA.Name, dogB.Name)
		}
		return utils.IntComparator(dogA.Age, dogB.Age)
	}

	tm := treeset.NewWith(dogComparator)

	for _, v := range testdata {
		tm.Add(v)
	}

	it := tm.Iterator()
	for it.Next() {
		v := it.Value()
		fmt.Printf("value: %v", v)
	}
}

func GodStringCmp() {
	fmt.Printf("res: %+v\n", utils.StringComparator("22fa1ba1-f395-4c7f-8b6c-3550f55734dd","a8d6d0b1-249c-4c0f-b82b-becbb9ca0872"))
}

func main(){
	//TestGods()
	GodStringCmp()
}