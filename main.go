package main

import (
	"fmt"
	"time"
)

func main() {
	startNow := time.Now()
	a := []int{1, 43, 65, 87, 12, 65, 8, 1, -1, 54, 0, 93, -12}

	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c

	fmt.Println("x:", x, ", y:", y)
	fmt.Println("took: ", time.Since(startNow))

}

func sum(a []int, c chan int) {

	sum := 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}

	c <- sum
}
