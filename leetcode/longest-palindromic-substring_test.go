package leetcode

import "testing"

var p5Cases = []([]string){
	[]string{"", ""},
	[]string{" ", " "},
	[]string{"b", "b"},
	[]string{"cbbd", "bb"},
	[]string{"bb", "bb"},
	[]string{"aaabaaaa", "aaabaaa"},
	[]string{"abcb", "bcb"},
	[]string{"ccc", "ccc"},
	[]string{"cccbbbaaaaaabbbbb", "bbbaaaaaabbb"},
	[]string{"cccbbbaaaaaaabbbbb", "bbbaaaaaaabbb"},
}

func TestP5v1(t *testing.T) {
	for _, v := range p5Cases {
		result := p5v1(v[0])
		if result == v[1] {
			t.Logf(checkMark)
		} else {
			t.Errorf("%s, get: %s, but should be: %s", ballotX, result, v[1])
		}
	}
}
