package stringsplosion

import (
	"fmt"
	"testing"
)

func ExampleStringSplosion() {
	fmt.Println(StringSplosion("Code"))
	// Output: CCoCodCode
}

func TestStringSplosion(t *testing.T) {
	tests := []struct {
		name string
		got  string
		want string
	}{
		{"Code", "Code", "CCoCodCode"},
		{"abc", "abc", "aababc"},
		{"empty", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringSplosion(tt.got); got != tt.want {
				t.Errorf("StringSplosion() = %v, want %v", got, tt.want)
			}
		})
	}
}
