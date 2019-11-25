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
		{"test case 1",1,"1"},
		{"test case 2",2,"2"},
		{"test case 3",3,"Fizz"},
		{"test case 4",4,"4"},
		{"test case 5",5,"Buzz"},
		{"test case 6",6,"Fizz"},
		{"test case 7",7,"7"},
		{"test case 8",8,"8"},
		{"test case 9",9,"Fizz"},
		{"test case 10",10,"Buzz"},
		{"test case 11",11,"11"},
		{"test case 12",12,"Fizz"},
		{"test case 13",13,"13"},
		{"test case 14",14,"14"},
		{"test case 15",15,"FizzBuzz"},
		{"test case 16",16,"16"},
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

/*
	for _, tt := range tests {
		if got := getFizzBuzz(tt.p); got != tt.want{
			t.Errorf("Want = %v, got = %v", tt.want, got)
		}
	}
*/

}