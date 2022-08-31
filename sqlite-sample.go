package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

var qry string = "select * from mydata where name like ?"

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
	nm := Input("name")
	ml := Input("mail")
	age := Input("age")
	ag, _ := strconv.Atoi(age)
	qry := "insert into mydata (name,mail,age) values (?,?,?)"
	con.Exec(qry, nm, ml, ag)
	showRecord(con)
}

func showRecord(con *sql.DB) {
	qry = "select * from mydata"
	rs, _ := con.Query(qry)
	for rs.Next() {
		fmt.Println(mydatafmRws(rs).Str())
	}
}

func mydatafmRws(rs *sql.Rows) *Mydata {
	var md Mydata
	err := rs.Scan(&md.ID, &md.Name, &md.Mail, &md.Age)
	if err != nil {
		panic(err)
	}
	return &md
}

func mydatafmRw(rs *sql.Row) *Mydata {
	var md Mydata
	err := rs.Scan(&md.ID, &md.Name, &md.Mail, &md.Age)
	if err != nil {
		panic(err)
	}
	return &md
}
