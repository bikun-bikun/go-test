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
				t.Errorf("Want = %d, got = %d", tt.want, got)
			}
		})
	}

}

func TestGetFizzBuzz(t *testing.T) {
	tests := []struct{
		name string
		p int
		want string
	}{
		{"Number",1,"1"},
		{"Fizz",3,"Fizz"},
		{"Buzz",5,"Buzz"},
		{"FizzBuzz",15,"FizzBuzz"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T){
			t.Parallel()
			if got := getFizzBuzz(tt.p); got != tt.want{
				t.Errorf("Want = %v, got = %v", tt.want, got)
			}
		})
	}

}