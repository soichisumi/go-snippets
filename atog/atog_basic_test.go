package atog

import (
	"testing"
	"github.com/yoyoyousei/SandBox/ext"
	"github.com/yoyoyousei/SandBox/testutil"
)

func TestGettingExtValue(t *testing.T) {
	actual1 := ext.Pi
	if !testutil.FloatEqE10(float64(3.141592), actual1) {
		t.Errorf(" %f が返されました", actual1)
	}

	actual2 := ext.GetPI()
	if !testutil.FloatEqE10(float64(3.14159265358979), actual2) {
		t.Errorf("%fが返されました", actual2)
	}
}

func TestAdd(t *testing.T) {
	actual := add(1, 2)
	if actual != 3 {
		t.Errorf("%dが返されました", actual)
	}
}

func TestAddByPointer(t *testing.T) {
	var x, y int = 1, 2
	addByPointer(&x, y)
	if x != 3 {
		t.Errorf("%dが返されました", x)
	}
}

func TestStruct(t *testing.T) {
	tmp2d := Point2d{x: 5, y: 10}
	tmp3d := Point3d{x: 10, y: 1, z: 111}

	actual2d := tmp2d.x
	actual3d := tmp3d.y

	if actual2d != 5 {
		t.Errorf("%dが返されました", actual2d)
	}

	if actual3d != 1 {
		t.Errorf("%dが返されました", actual3d)
	}
}

func TestArrayAndSlice(t *testing.T) {
	var a = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	if a[0] != 1 {
		t.Errorf("err")
	}
	s:=a[len(a)/2:] //スライスを作成。スライスは配列への参照
	if s[0] != 6{
		t.Errorf("er")
	}
}
