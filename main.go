package main

import (
	"fmt"

	"github.com/gingama4/bustatus/config"
)

func main() {
	c, _ := config.GetConfig("")
	fmt.Printf("%+v\n", c)
}
