package repl

import (
	"bufio"
	"fmt"
	"io"
	"writing/lexer"
	"writing/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}
		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}

const DOG_FACE = `
  / \__
 (    @\__ 
 /         O
/   (_____/
/_____/   U
`

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, DOG_FACE)
	io.WriteString(out, "Oops! Looks like our pup got paws-deep in code!\n")
	io.WriteString(out, "Fetching errors... Please wait...\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
