package main

import (
	"douyin/biz/pack"
	"fmt"
)

func main() {
	token, err := pack.MakeToken(207255505750986753, "wqrfasd")
	if err != nil {
		return
	}
	fmt.Println(token)
}
