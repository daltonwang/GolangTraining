package main

import "fmt"

const metersToYards float64 = 1.09361

func main() {
	var meters float64
	fmt.Print("Enter meters swam: ")
	fmt.Scan(&meters) //Scan 里面的参数需要是变量的地址
	yards := meters * metersToYards
	fmt.Println(meters, " meters is ", yards, " yards.")
}
