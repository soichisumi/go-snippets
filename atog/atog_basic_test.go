package atog

import(
	"testing"
	"github.com/yoyoyousei/SandBox/ext"
	"github.com/yoyoyousei/SandBox/testutil"
)

func TestGettingExtValue(t *testing.T){
	actual1 := ext.Pi
	if !testutil.FloatEqE10(float64(3.141592), actual1){
		t.Errorf(" %f が返されました", actual1)
	}

	actual2 := ext.GetPI()
	if !testutil.FloatEqE10(float64(3.14159265358979), actual2){
		t.Errorf("%fが返されました", actual2)
	}
}

func TestAdd(t *testing.T){
	actual := add(1,2)
	if actual != 3{
		t.Errorf("%dが返されました")
	}
}