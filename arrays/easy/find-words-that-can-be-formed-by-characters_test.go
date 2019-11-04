package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs6 is testcase slice
var tcs6 = []struct {
	words []string
	chars string
	ans   int
}{

	{
		[]string{"cat", "bt", "hat", "tree"},
		"atach",
		6,
	},

	{
		[]string{"hello", "world", "leetcode"},
		"welldonehoneyr",
		10,
	},

	// 可以有多个 testcase
}

func Test_countCharacters(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs6 {
		ast.Equal(tc.ans, countCharacters(tc.words, tc.chars), "输入:%v", tc)
	}
}

func Benchmark_countCharacters(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs6 {
			countCharacters(tc.words, tc.chars)
		}
	}
}
