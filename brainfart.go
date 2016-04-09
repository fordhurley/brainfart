package brainfart

import (
	"io"
	"log"
)

// Run runs the program.
func Run(program []byte, input io.ByteReader, output io.ByteWriter) {
	var i int       // instruction pointer
	var marks []int // loop start markers

	var d int // data pointer
	var data = make([]byte, 1024)

	for i = 0; i < len(program); i++ {
		var err error
		switch program[i] {
		case '>':
			d++
		case '<':
			d--
		case '+':
			data[d]++
		case '-':
			data[d]--
		case '.':
			err = output.WriteByte(data[d])
		case ',':
			var b byte
			b, err = input.ReadByte()
			if err == io.EOF {
				return
			}
			data[d] = b
		case '[':
			if data[d] == 0 {
				// Jump forward to end of loop:
				depth := 1
				for depth > 0 {
					i++
					switch program[i] {
					case '[':
						// Nested loop, so skip the next closing bracket:
						depth++
					case ']':
						depth--
					}
				}
			} else {
				// Remember this so we can jump back:
				marks = append(marks, i)
			}
		case ']':
			if data[d] != 0 {
				// Jump back to beginning of loop:
				i = marks[len(marks)-1]
			} else {
				// We're done with this loop, so forget the matching mark:
				marks = marks[:len(marks)-1]
			}
		}
		if err != nil {
			log.Fatal(err)
		}
	}
}
