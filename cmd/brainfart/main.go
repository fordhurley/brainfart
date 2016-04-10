package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fordhurley/brainfart"
)

func main() {
	var program []byte
	var err error

	if len(os.Args) < 2 {
		// Read program from stdin
		program, err = ioutil.ReadAll(os.Stdin)
	} else {
		// Read program from file named by first argument
		filename := os.Args[1]
		program, err = ioutil.ReadFile(filename)
	}

	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	brainfart.Run(program, in, out)
	out.Flush()
}
