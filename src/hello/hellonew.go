package main

import (
	"fmt"
	"time"
	"math/rand"
	"math"
)

//func log(message string) {
//}
//
//func add(a int, b int) int {
//	return a + b
//}
//
//func power(name string) (int, bool) {
//	return 1, true
//}

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, python, java bool


func main() {
	fmt.Println("The time is", time.Now())
	fmt.Println("My favourite number is ", rand.Intn(10))
	fmt.Printf("You have %g problems", math.Sqrt(7))
	fmt.Println(math.Pi)
	fmt.Println(add(100, 200))
	a, b := swap("Shashank", "Vivek")
	fmt.Println(a, b)
	fmt.Println(split(17))
	//var i int
	//var c, p, jav = false, false, "Hey"
	//fmt.Println(i, c, java, python)
	//fmt.Println(c, p, jav)
	//var i int = 42
	//var f float64 = float64(i)
	//var u uint = uint(f)
	//fmt.Println(i, f, u)
	//const age = 100
	//fmt.Println(age)
	//fmt.Println(age)
	//me := 3
	//fmt.Println(me)
}