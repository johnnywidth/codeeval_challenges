package main

import "testing"

func TestLCS(t *testing.T) {
	l := "XMJYAUZ"
	r := "MZJAWXU"
	expect := "MJAU"

	res := LCS(l, r)
	if res != expect {
		t.Fatalf("`%s` not equal with `%s`", res, expect)
	}

	l = "thisisatest"
	r = "testing123testing"
	expect = "tsitest"

	res = LCS(l, r)
	if res != expect {
		t.Fatalf("`%s` not equal with `%s`", res, expect)
	}
}

func BenchmarkLCS(b *testing.B) {
	l := "XMJYAUZ"
	r := "MZJAWXU"
	for i := 0; i < b.N; i++ {
		LCS(l, r)
	}
}
