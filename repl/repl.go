package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/ayushbhargav/monkey-lang/token"

	"github.com/ayushbhargav/monkey-lang/lexer"
)

// PROMPT for each command in REPL
const PROMPT = ">> "

// Start REPL
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
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}

	}

}
