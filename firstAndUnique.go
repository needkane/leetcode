/* Return the first unique byte of a string

example :

"aabc" - > b

"aabcb" -> c

*/

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type entry struct {
	times int
	pos   int
}

func firstUnique(str string) (result rune) {

	rs := []rune(str)
	if len(rs) == 0 {
		return
	}
	m0 := make(map[rune]*entry) //times and pos
	m1 := make(map[rune]int)    //only unique
	for i := 0; i < len(rs); i++ {

		value := rs[i]
		v, ok := m0[value]
		if ok {
			m0[value].times = v.times + 1
		} else {
			m0[value] = &entry{1, i}

		}
	}
	for k, v := range m0 {
		if v.times == 1 {
			m1[k] = v.pos
		}
	}

	index := len(rs) - 1
	for _, v := range m1 {
		if v < index {
			index = v
		}
	}
	result = rs[index]
	return
}

func main() {
	var t *testing.T
	assert.Equal(t, firstUnique("aabc"), 'b')
	assert.Equal(t, firstUnique("aabcb"), 'c')
	assert.Equal(t, firstUnique("cddaabc"), 'b')
}
