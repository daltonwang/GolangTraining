package main

import "fmt"

func main() {
	for i := 60; i < 122; i++ {
		//%d 10进制
		//%b 2进制
		//%x 16进制
		//%#x 0x前缀的16进制
		//%#q utf-8 字符
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	}
}
