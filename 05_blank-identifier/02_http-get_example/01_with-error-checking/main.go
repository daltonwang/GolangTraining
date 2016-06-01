package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//http.Get 用来获取网页
	res, err := http.Get("http://www.mcleods.com/")
	if err != nil {
		log.Fatal(err)
	}
	
	//ioutil.ReadAll用来解析网页
	page, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", page)
}
