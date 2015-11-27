package main

import (
	"fmt"
	"time"
	"math"
	)

func add(x int, y int) int{
	return x+y
	}

func swap(x string, y string) (string, string) {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += 1
		}
	fmt.Println( sum )
	return y, x
	}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x,n); v<lim{
		return v
		}
	return lim
}

func main() {
	fmt.Println("Hello");
	fmt.Println("Current Time ", time.Now());
	fmt.Println(add(42,13));
	a,b := swap("hello", "world")
	fmt.Println(a,b);
	fmt.Println(pow(3,2,10));
}