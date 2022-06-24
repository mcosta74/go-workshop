package utils

import (
	"testing"
	"unicode/utf8"
)

func TestAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"(0,1)", args{0, 1}, 1},
		{"(1,1)", args{1, 1}, 2},
		{"(2,1)", args{2, 1}, 3},
		{"(1,-1)", args{1, -1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"both positive", args{1, 1}, 1, false},
		{"both negative", args{-1, -1}, 1, false},
		{"negative result", args{1, -1}, -1, false},
		{"division by zero", args{1, 0}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Divide(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Divide() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Divide() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"hello world", "dlrow olleh"},
		{" ", " "},
		{"12345!", "!54321"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := Reverse(tt.input); got != tt.want {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkReverse(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Reverse("Hello, World!")
	}
}

func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev := Reverse(orig)
		doubleRev := Reverse(rev)
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}

// rev := ReverseV2(orig)
// doubleRev := ReverseV2(rev)
// if orig != doubleRev {
// 	t.Errorf("Before: %q, after: %q", orig, doubleRev)
// }
// if utf8.ValidString(orig) && !utf8.ValidString(rev) {
// 	t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
// }

// 	rev, revErr := ReverseV3(orig)
// 	if revErr != nil {
// 		return
// 	}
//
// 	doubleRev, doubleRevErr := ReverseV3(rev)
// 	if doubleRevErr != nil {
// 		return
// 	}
// 	if orig != doubleRev {
// 		t.Errorf("Before: %q, after: %q", orig, doubleRev)
// 	}
// 	if utf8.ValidString(orig) && !utf8.ValidString(rev) {
// 		t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
// 	}
