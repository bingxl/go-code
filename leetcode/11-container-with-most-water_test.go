package leetcode

import "testing"
var p11Cases = []struct{
	c []int;
	r int;
} {
	{[]int{1,8,6,2,5,4,8,3,7}, 49},
	{[]int{10,2,3,4,5,1,6,10}, 70},

}

func p11Base(t *testing.T, f func([]int)int) {
	for _, c := range p11Cases {
		r := f(c.c)
		if r != c.r {
			t.Errorf("%s, string: %v, except: %v, but got:%v", ballotX, c.c, c.r, r)
		} else {
			t.Logf(checkMark)
		}
	}
}
func TestP11(t *testing.T) {
	t.Run("p11v1", func (t *testing.T) {p11Base(t, p11v1)})
	t.Run("p11v2", func (t *testing.T) {p11Base(t, p11v2)})
}