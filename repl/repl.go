package repl

import (
	"bufio"
	"fmt"
	"io"
	"github.com/sqqueak/monkey/lexer"
	"github.com/sqqueak/monkey/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return // scanned nothing => end parsing this line
		}

		line := scanner.Text()
		l := lexer.New(line)

		// getting the internal Token representation of input
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
