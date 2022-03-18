package main

import (
	"fmt"
)

func main() {
	fmt.Println("-----------")
	//p := flag.Bool("p", false, "default path")
	//flag.Parse()

	var load LoadConfigurations
	config := load.LoadConfig()
	fmt.Println(len(config.Server.Env))
}
