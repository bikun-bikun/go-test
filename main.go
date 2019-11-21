package main

import (
	"strconv"
	"time"
)

func main(){

}

func Sum(a, b int) (s int){
	s = a + b
	return
}

func getFizzBuzz(p int) (r string){
	if p % 3 == 0 {
		r = r + "Fizz"
	}

	if p % 5 == 0 {
		r = r + "Buzz"
	}

	if r == "" {
		r = strconv.Itoa(p)
	}

	time.Sleep(100 * time.Millisecond)

	return
}