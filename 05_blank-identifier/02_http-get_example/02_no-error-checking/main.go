package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//这里下面两个函数直接省略的返回的ERR参数，不做参数校验
	res, _ := http.Get("http://www.mcleods.com/")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
}
