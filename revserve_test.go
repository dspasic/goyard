package main

import (
	"testing"
	"unicode/utf8"
)

func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, World", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev, err := Reverse(orig)
		if err != nil {
			return
		}
		doubleRev, doubleRevErr := Reverse(rev)
		if doubleRevErr != nil {
			return
		}
		if doubleRev != orig {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produce invalid utf-8 stringL: %q", rev)
		}

	})
}
