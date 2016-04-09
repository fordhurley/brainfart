package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fordhurley/brainfart"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <filename>", os.Args[0])
		os.Exit(1)
	}

	program, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	brainfart.Run(program, in, out)
	out.Flush()
}
