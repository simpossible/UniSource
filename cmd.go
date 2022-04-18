package main

import (
	"cmd.unisource.icu/config"
	"fmt"
)

func main() {
	unis, _ := config.AllConfigAtDir("/Users/liangjinfeng/dev/unisource/cmd.unisource.icu/test/test_data")
	fmt.Println("all unis is %v", unis)
}
