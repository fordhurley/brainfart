package brainfart

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func checkOutput(t *testing.T, program string, input []byte, expected []byte) {
	var in = bytes.NewReader(input)
	var out = new(bytes.Buffer)
	Run([]byte(program), in, out)
	assert.Equal(t, out.Bytes(), expected)
}

func TestOutput(t *testing.T) {
	checkOutput(t,
		"...",
		[]byte{},
		[]byte{0, 0, 0},
	)
}

func TestIncrementDecrementCellValue(t *testing.T) {
	checkOutput(t,
		".+.+.+.-.-.-.",
		[]byte{},
		[]byte{0, 1, 2, 3, 2, 1, 0},
	)
}

func TestIncrementDecrementInstrPointer(t *testing.T) {
	checkOutput(t,
		".>+.>++.>+++.<.<.<.",
		[]byte{},
		[]byte{0, 1, 2, 3, 2, 1, 0},
	)
}

func TestAddCells(t *testing.T) {
	// 5 + 3 = 8
	checkOutput(t,
		"+++++.>+++.<[->+<].>.",
		[]byte{},
		[]byte{5, 3, 0, 8},
	)
}

func TestInput(t *testing.T) {
	// Input two values and add them together
	checkOutput(t,
		",.>,.<[->+<]>.",
		[]byte{42, 11},
		[]byte{42, 11, 53},
	)
}

func TestNestedLoops(t *testing.T) {
	checkOutput(t,
		".[.[.++++.].].",
		[]byte{},
		[]byte{0, 0},
	)
}

var helloWorld = `++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.`
var rot13 = `
	-,+[                         Read first character and start outer character reading loop
			-[                       Skip forward if character is 0
					>>++++[>++++++++<-]  Set up divisor (32) for division loop
																 (MEMORY LAYOUT: dividend copy remainder divisor quotient zero zero)
					<+<-[                Set up dividend (x minus 1) and enter division loop
							>+>+>-[>>>]      Increase copy and remainder / reduce divisor / Normal case: skip forward
							<[[>+<-]>>+>]    Special case: move remainder back to divisor and increase quotient
							<<<<<-           Decrement dividend
					]                    End division loop
			]>>>[-]+                 End skip loop; zero former divisor and reuse space for a flag
			>--[-[<->+++[-]]]<[         Zero that flag unless quotient was 2 or 3; zero quotient; check flag
					++++++++++++<[       If flag then set up divisor (13) for second division loop
																 (MEMORY LAYOUT: zero copy dividend divisor remainder quotient zero zero)
							>-[>+>>]         Reduce divisor; Normal case: increase remainder
							>[+[<+>-]>+>>]   Special case: increase remainder / move it back to divisor / increase quotient
							<<<<<-           Decrease dividend
					]                    End division loop
					>>[<+>-]             Add remainder back to divisor to get a useful 13
					>[                   Skip forward if quotient was 0
							-[               Decrement quotient and skip forward if quotient was 1
									-<<[-]>>     Zero quotient and divisor if quotient was 2
							]<<[<<->>-]>>    Zero divisor and subtract 13 from copy if quotient was 1
					]<<[<<+>>-]          Zero divisor and add 13 to copy if quotient was 0
			]                        End outer skip loop (jump to here if ((character minus 1)/32) was not 2 or 3)
			<[-]                     Clear remainder from first division if second division was skipped
			<.[-]                    Output ROT13ed character from copy and clear it
			<-,+                     Read next character
	]                            End character reading loop
`

func TestHelloWorld(t *testing.T) {
	checkOutput(t,
		helloWorld,
		[]byte{},
		[]byte("Hello World!\n"),
	)
}

func ExampleRun_rot13() {
	program := []byte(rot13)
	var input = bytes.NewReader([]byte("AaBbCc XxYyZz 123"))
	var output = new(bytes.Buffer)
	Run(program, input, output)
	fmt.Println(output.String())

	// Output:
	// NnOoPp KkLlMm 123
}

func BenchmarkHelloWorld(b *testing.B) {
	program := []byte(helloWorld)
	var input = new(bytes.Buffer)
	for i := 0; i < b.N; i++ {
		var output = new(bytes.Buffer)
		Run(program, input, output)
	}
}

func BenchmarkRot13(b *testing.B) {
	program := []byte(rot13)
	in := []byte("abcdefghijklmnopqrstuvwxyz1234567890")
	for i := 0; i < b.N; i++ {
		var input = bytes.NewReader(in)
		var output = new(bytes.Buffer)
		Run(program, input, output)
	}
}
