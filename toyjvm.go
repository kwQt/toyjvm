package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no class file path")
		return
	}
	path := os.Args[1]

	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	parseClassFile(data)
}

func parseClassFile(data []byte) {
	fmt.Println(hex.Dump(data))
}
