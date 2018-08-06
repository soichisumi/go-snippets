package golang

import "fmt"

var c, python, java bool = true, false, true

type Point2d struct {
	x, y int
}

type Point3d struct {
	x, y, z int
}

func useVariables() {
	fmt.Sprintf("%t %t %t", c, python, java)
}

func addByPointer(p *int, num int) {
	*p += num
}

func add(x int, y int) int {
	return x + y
}
