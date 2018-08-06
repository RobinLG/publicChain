package main

import (
	"os"
	"fmt"
)

func main() {

	args := os.Args

	fmt.Printf("%v\n", args)
	fmt.Printf("%v\n", args[1])

}

//main -open -printchain "aaa" -number 9