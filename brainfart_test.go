package brainfart

import (
	"bytes"
	"fmt"
)

// Using examples here just because it's easier to check the output.

func ExampleRun() {
	// Output
	inst := []byte("...")
	var input = new(bytes.Reader)
	var output = new(bytes.Buffer)
	Run(inst, input, output)
	fmt.Println(output.Bytes())

	// Increment/decrement cell value
	inst = []byte(".+.+.+.-.-.-.")
	output.Reset()
	Run(inst, input, output)
	fmt.Println(output.Bytes())

	// Increment/decrement instruction pointer
	inst = []byte(".>+.>++.>+++.<.<.<.")
	output.Reset()
	Run(inst, input, output)
	fmt.Println(output.Bytes())

	// Add two cells (5 + 3 = 8)
	inst = []byte("+++++.>+++.<[->+<].>.")
	output.Reset()
	Run(inst, input, output)
	fmt.Println(output.Bytes())

	// Input two values and add them together
	inst = []byte(",.>,.<[->+<]>.")
	input = bytes.NewReader([]byte{42, 11})
	output.Reset()
	Run(inst, input, output)
	fmt.Println(output.Bytes())

	// Output:
	// [0 0 0]
	// [0 1 2 3 2 1 0]
	// [0 1 2 3 2 1 0]
	// [5 3 0 8]
	// [42 11 53]
}

func ExampleRun_helloworld() {
	inst := []byte(`++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.`)
	var input = new(bytes.Reader)
	var output = new(bytes.Buffer)
	Run(inst, input, output)
	fmt.Println(output.String())

	// Output:
	// Hello World!
}
