package main

import "testing"

func TestSum(t *testing.T) {
	tests := []struct{
		name string
		a, b, want int
	}{
		{"Simple", 0, 1, 1},
		{"Minus", -1 , -1, -2},
		{"Both", -3, 2, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			if got := Sum(tt.a, tt.b); got != tt.want{
				t.Fatalf("Want = %d, got = %d", tt.want, got)
			}
		})
	}

}