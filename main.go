package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	fmt.Println("testtest", "test")
	resp, err := http.Get("https://google.com")
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("experienced error:",err)
	}else{
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(body))
		}
	}
}
