package main

import (
	"flag"
	"fmt"

	"github.com/youthke/go-crud/conf"
	"github.com/youthke/go-crud/pkg/server"
)

var state = flag.String("s", "local", "local/prd")

func main() {

	flag.Parse()
	err := conf.SetUp(fmt.Sprintf("conf/env/%s.toml", *state))

	if err != nil{
		fmt.Printf("%s", err)
		return
	}

	conf.Init()
	server.Init()
}
