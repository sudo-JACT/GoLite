package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

func printquery(rows *sql.Rows) {

	col, _ := rows.Columns()

	n := len(col)

	colums := make([]any, n)

	for i := 0; i < n; i++ {

		var tmp any

		colums[i] = &tmp

	}

	for rows.Next() {

		var s string = ""

		rows.Scan(colums...)

		for _, c := range colums {

			s += fmt.Sprintf("%v ", *(c.(*any)))

		}

		fmt.Println(s)

	}

}

func scanner(query *string) {

	scan := bufio.NewScanner(os.Stdin)

	scan.Scan()

	*query = scan.Text()

}

func main() {

	var err error
	var rows *sql.Rows
	var db *sql.DB
	var q *sql.Stmt
	var flag bool = true
	var query string

	defer func() {

		if r := recover(); r == nil {

			db.Close()
			rows.Close()

		}

	}()

	fmt.Print("filename/filepath >> ")
	scanner(&query)

	db, err = sql.Open("sqlite", query)

	if err != nil {

		fmt.Println(err)

	}

	for flag {

		fmt.Print(">> ")

		scanner(&query)

		if query == "quit" {

			flag = false

		} else {

			q, err = db.Prepare(query)

			if err == nil {

				rows, err = q.Query()

				printquery(rows)

			} else {

				fmt.Println(err)

			}
		}

	}
}
