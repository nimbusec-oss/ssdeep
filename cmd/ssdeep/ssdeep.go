package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/nimbusec-oss/ssdeep"
)

func main() {
	flag.Parse()
	if len(flag.Args()) < 1 {
		fmt.Println("Please provide a file path: ./ssdeep /tmp/file")
		return
	}
	if len(flag.Args()) == 1 {
		file, err := os.Open(flag.Args()[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		hash, err := ssdeep.Fuzzy(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(hash.String())
	}
	if len(flag.Args()) == 2 {
		f1, err := os.Open(flag.Args()[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		h1, err := ssdeep.Fuzzy(f1)
		if err != nil {
			fmt.Println(err)
			return
		}
		f2, err := os.Open(flag.Args()[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		h2, err := ssdeep.Fuzzy(f2)
		if err != nil {
			fmt.Println(err)
			return
		}
		score := ssdeep.HashDistance(h1, h2)
		fmt.Println(score)
	}
}
