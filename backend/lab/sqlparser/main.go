package main

import (
	"io"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/xwb1989/sqlparser"
)

func main() {
	r := strings.NewReader("INSERT INTO table1 VALUES (1, 'a'); INSERT INTO table2 VALUES (3, 4);")

	tokens := sqlparser.NewTokenizer(r)
	for {
		stmt, err := sqlparser.ParseNext(tokens)
		if err == io.EOF {
			break
		}
		spew.Dump(stmt)
		// Do something with stmt or err.
	}
}
