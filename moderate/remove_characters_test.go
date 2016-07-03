package main

import (
	"testing"
)

func TestRemoveCharacters(t *testing.T) {
	s := "how are you"
	e := "abc"
	expect := "how re you"

	r := RemoveCharacters(s, e)
	if r != expect {
		t.Fatalf("`%s` not equal with `%s`", r, expect)
	}
}

func BenchmarkRemoveCharacters(b *testing.B) {
	s := "Replace returns a copy of the string s with the first n non-overlapping instances of old replaced by new. If old is empty, it matches at the beginning of the string and after each UTF-8 sequence, yielding up to k+1 replacements for a k-rune string. If n < 0, there is no limit on the number of replacements."
	e := "abcwfdshkjngl"
	for i := 0; i < b.N; i++ {
		RemoveCharacters(s, e)
	}
}
