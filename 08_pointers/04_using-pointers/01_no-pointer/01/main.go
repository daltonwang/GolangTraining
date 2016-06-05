package main

import "fmt"

func zero(z int) {
	z = 0 //进入子函数后，如果需要更改这个变量，则需要引用这个变量的指针
}

func main() {
	x := 5
	zero(x) /
	fmt.Println(x) // x is still 5
}
