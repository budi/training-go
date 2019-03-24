package longestsubsequence

import (
	"fmt"
	"testing"
)

func Test_byLength_Len(t *testing.T) {
	tests := []struct {
		name string
		s    byLength
		want int
	}{
		{"one", byLength{"s"}, 1},
		{"three", byLength{"s", "t", "r"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Len(); got != tt.want {
				t.Errorf("byLength.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_byLength_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		s    byLength
		args args
		want byLength
	}{
		{"one to two", byLength{"s", "t", "r"}, args{0, 1}, byLength{"t", "s", "r"}},
		{"one to three", byLength{"s", "t", "r"}, args{0, 2}, byLength{"r", "t", "s"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := tt.s[tt.args.i]
			j := tt.s[tt.args.j]
			tt.s.Swap(tt.args.i, tt.args.j)
			if tt.s[tt.args.j] != i || tt.s[tt.args.i] != j {
				t.Errorf("byLength.Swap() = %v, want %v", tt.s, tt.want)
			}
		})
	}
}

func Test_byLength_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		s    byLength
		args args
		want bool
	}{
		{
			"true",
			byLength{"loooooong", "short"},
			args{0, 1},
			true,
		},
		{
			"false",
			byLength{"short", "loooooong"},
			args{0, 1},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("byLength.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleIsSubsequence_false() {
	fmt.Println(IsSubsequence("abppplee", "kangaroo"))
	// Output: false
}

func ExampleIsSubsequence_true() {
	fmt.Println(IsSubsequence("abppplee", "able"))
	// Output: true
}

func Test_isSubsequence(t *testing.T) {
	type args struct {
		ref  string
		word string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"success 1",
			args{"abppplee", "able"},
			true,
		},
		{
			"success 2",
			args{"abppplee", "ale"},
			true,
		},
		{
			"success 3",
			args{"abppplee", "apple"},
			true,
		},
		{
			"fail 1",
			args{"abppplee", "bale"},
			false,
		},
		{
			"fail 2",
			args{"abppplee", "kangaroo"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sub := IsSubsequence(tt.args.ref, tt.args.word)
			if sub != tt.want {
				t.Fatalf("Unexpected value: want %v, got %v", tt.want, sub)
			}
		})
	}
}

func ExampleFindLongestSubsequence_apple() {
	fmt.Println(FindLongestSubsequence("abppplee", []string{"able", "ale", "apple"}))
	// Output: apple
}

func ExampleFindLongestSubsequence_empty() {
	fmt.Printf("\"%v\"", FindLongestSubsequence("abppplee", []string{"kangaroo", "bird", "love"}))
	// Output: ""
}

func Test_findLongestSub(t *testing.T) {
	type args struct {
		ref   string
		words []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"apple",
			args{
				"abppplee",
				[]string{"able", "ale", "apple"},
			},
			"apple",
		},
		{
			"empty",
			args{
				"abppplee",
				[]string{"kangaroo", "bird", "love"},
			},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindLongestSubsequence(tt.args.ref, tt.args.words); got != tt.want {
				t.Errorf("FindLongestSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
