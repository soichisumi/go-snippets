package main

import (
	"fmt"
	"regexp"
)

func main() {
	r, _ := regexp.Compile("p([a-z]+)ch")
	fmt.Printf("FindAllString: %+v\n", r.FindAllString("peach-punch", -1))
	fmt.Printf("FindStringSubmatch: %+v\n", r.FindStringSubmatch("peach-punch"))
	fmt.Printf("FindStringSubmatch: %+v\n", r.FindAllStringSubmatch("peach-punch", -1))
}
