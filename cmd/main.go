package main

import (
	"bytes"
	"fmt"

	"github.com/eatonphil/gosql"
)

func main() {
	mb := gosql.NewMemoryBackend()

	source := bytes.NewBufferString("CREATE TABLE users (id INT, name TEXT); INSERT INTO users VALUES (1, 'Admin'); SELECT id, name FROM users")

	ast, err := gosql.Parse(source)
	if err != nil {
		panic(err)
	}

	err = mb.CreateTable(ast.Statements[0].CreateTableStatement)
	if err != nil {
		panic(err)
	}

	err = mb.Insert(ast.Statements[1].InsertStatement)
	if err != nil {
		panic(err)
	}

	results, err := mb.Select(ast.Statements[2].SelectStatement)
	if err != nil {
		panic(err)
	}

	for _, col := range results.Columns {
		fmt.Printf("| %s ", col.Name)
	}
	fmt.Println()

	for _, result := range results.Rows {
		fmt.Printf("| ")

		for i, cell := range result {
			typ := results.Columns[i].Type
			s := ""
			switch typ {
			case gosql.IntType:
				s = fmt.Sprintf("%d", cell.AsInt())
			case gosql.TextType:
				s = cell.AsText()
			}

			fmt.Printf(" %s | ", s)
		}

		fmt.Println()
	}
}