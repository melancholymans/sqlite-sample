package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type Mydata struct {
	ID   int
	Name string
	Mail string
	Age  int
}

func (m *Mydata) Str() string {
	return "<\"" + strconv.Itoa(m.ID) + ":" + m.Name + "\" " + m.Mail + "," + strconv.Itoa(m.Age) + ">"
}

func main() {
	con, err := sql.Open("sqlite3", "data.db3")
	if err != nil {
		panic(err)
	}
	defer con.Close()
	q := "select * from mydata"
	rs, err := con.Query(q)
	if err != nil {
		panic(err)
	}
	for rs.Next() {
		var md Mydata
		err := rs.Scan(&md.ID, &md.Name, &md.Mail, &md.Age)
		if err != nil {
			panic(err)
		}
		fmt.Println(md.Str())
	}
}
