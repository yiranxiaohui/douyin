package main

import (
	"douyin/biz/pack"
	"fmt"
)

func main() {
	token, err := pack.MakeToken(1, "12324255")
	if err != nil {
		return
	}
	fmt.Println(token)
}
