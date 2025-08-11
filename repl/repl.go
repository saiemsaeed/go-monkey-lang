package repl

import (
	"bufio"
	"fmt"
	"os"

	"github.com/saiemsaeed/monkey-go/lexer"
	"github.com/saiemsaeed/monkey-go/token"
)

const PROMT = ">> "

func Start() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println(PROMT)
		scan := scanner.Scan()

		if !scan {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
