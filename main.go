package main

import (
	"dynamic_json/parser"
	"fmt"
	"os"
)

func main() {
	var data []byte
	var p *parser.Parser
	var err error

	if data, err = os.ReadFile("config.json"); err != nil {
		panic(err)
	}

	p, err = parser.Init(data)
	if err != nil {
		panic(err)
	}

	if data, err = p.ToFront(); err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
