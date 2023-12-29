package main

import (
	"bytes"
	"testing"
)

func TestSolve(t *testing.T) {
	want := `
		517 946 283
		348 512 976
		926 837 154

		295 763 418
		163 498 725
		784 125 639

		471 359 862
		832 674 591
		659 281 347
	`
	s := `
		..7 .4. 28.
		..8 .12 9..
		.2. ... 1..

		... 76. ...
		16. ... ..5
		.8. ... ...

		.7. ... .6.
		... ..4 ..1
		..9 .81 3.7
	`
	onUpdate := func(board [][]byte) {
		// do nothing
	}
	got := Solve(StringToBoard(s), onUpdate)
	if !bytesArrayEqual(got, StringToBoard(want)) {
		t.Errorf("Solve(%v)\ngot:\n%v\nwant:%v", s, bytesArrayToString(got), want)
	}
}

func TestStringToBoard(t *testing.T) {
	want := [][]byte{
		{46, 46, 55, 46, 52, 46, 50, 56, 46},
		{46, 46, 56, 46, 49, 50, 57, 46, 46},
	}
	s := `
		..7 .4. 28.

		..8 .12 9..
	`
	got := StringToBoard(s)
	if !bytesArrayEqual(got, want) {
		t.Errorf("StringToBoard(%v)\n got %v\nwant %v", s, got, want)
	}
}

func bytesArrayEqual(a, b [][]byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !bytes.Equal(a[i], b[i]) {
			return false
		}
	}
	return true
}

func bytesArrayToString(a [][]byte) string {
	var s string
	for _, b := range a {
		s += string(b) + "\n"
	}
	return s
}
