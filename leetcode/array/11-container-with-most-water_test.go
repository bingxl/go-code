package array

import "testing"

func TestP11(t *testing.T) {
	var p11Cases = []struct {
		c []int
		r int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{10, 2, 3, 4, 5, 1, 6, 10}, 70},
	}

	methods := map[string]func([]int) int{
		"p11MaxAreaV1": p11MaxAreaV1,
		"p11MaxAreaV2": p11MaxAreaV2,
	}
	for name, method := range methods {
		t.Run(name, func(t *testing.T) {
			for _, c := range p11Cases {
				r := method(c.c)
				if r != c.r {
					t.Errorf("%s, string: %v, except: %v, but got:%v", "error", c.c, c.r, r)
				} else {
					t.Logf("success")
				}
			}
		})
	}

}
