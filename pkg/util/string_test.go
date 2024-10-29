package util

import (
	"testing"
)

func TestRandStrGen(t *testing.T) {
	strLen := 10
	numLoop := 5000

	for i := 0; i < numLoop; i++ {
		str := RandString(strLen)

		if len(str) != strLen {
			t.Fatalf(`Output string doesn't have the required length`)
		}

		for _, char := range str {
			if (char < rune('A') || char > rune('F')) && (char < rune('0') || char > rune('9')) {
				t.Fatalf(`Character:%s is out of bound `, string(char))
			}

		}
	}

}

func BenchmarkRandStrGen(b *testing.B) {
	strLen := 10
	for n := 0; n < b.N; n++ {
		_ = RandString(strLen)

	}
}
