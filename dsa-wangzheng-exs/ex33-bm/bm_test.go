package ex33_bm

import (
	"fmt"
	"testing"
)

var tests = []struct{
	str string
	b string
	index int
}{
	{"abcacabdc", "abd", 5},
	{"aaaaaaaaaaaaa", "aa", 0},
	{"aaaaaaa", "a", 0},
	{"aaaaaaaaaa", "ba", -1},	// 这种情况下模式串按照坏字符规则向左滑动了，引起索引越界，panic
}

// 仅采用坏字符规则的BM，效率极高，但是部分情况会出现模式串向左滑动的情况，所以仍需加上好后缀规则。
func TestBMOnlyBadChar(t *testing.T) {
	for _, test := range  tests {
		i := BMOnlyBadChar(test.str, test.b)
		if i != test.index {
			t.Errorf("str=%s, b=%s, index should be %d, but get %d", test.str, test.b, test.index, i)
		}
	}
}

func TestGenerateGS(t *testing.T) {
	suffix, prefix := generateGS("cabcab")
	fmt.Printf("suffix= %#v\n", suffix)
	fmt.Printf("prefix= %#v\n", prefix)
}

func TestBM(t *testing.T) {
	for _, test := range  tests {
		i := BM(test.str, test.b)
		if i != test.index {
			t.Errorf("str=%s, b=%s, index should be %d, but get %d", test.str, test.b, test.index, i)
		}
	}
}

func TestBMOnlyGoodSuffix(t *testing.T) {
	for _, test := range  tests {
		i := BMOnlyGoodSuffix(test.str, test.b)
		if i != test.index {
			t.Errorf("str=%s, b=%s, index should be %d, but get %d", test.str, test.b, test.index, i)
		}
	}
}