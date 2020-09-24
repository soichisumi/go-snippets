package locking

import (
	"github.com/google/go-cmp/cmp"
	"sync/atomic"
	"testing"
)

func Test_TryLock(t *testing.T) {
	tests := []struct {
		name string
		locker uint32
		res []bool
	}{
		{
			name: "success",
			locker: 0,
			res: []bool{true, false, false, false, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := make([]bool, 0, 5)
			for i:= 0; i<4; i++ {
				res = append(res, atomic.CompareAndSwapUint32(&tt.locker, 0, 1)) // trylock
			}
			atomic.StoreUint32(&tt.locker, 0) //free
			res = append(res, atomic.CompareAndSwapUint32(&tt.locker, 0, 1))

			if !cmp.Equal(res, tt.res) {
				t.Errorf("fail, res: %+v, exp: %+v\n", res, tt.res)
				return
			}
		})
	}
}