package server

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	protos "github.com/akashkumar8/micproject/protos"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hashicorp/go-hclog"
)

// Course is
type Course struct {
	log hclog.Logger
}

var (
	name   string
	value  string
	count  int32
	hour   string
	repeat int32
	id     int
	date   string
	result string
)

var db *sql.DB
var err error
var dhour int

// NewCourse is
func NewCourse(l hclog.Logger) *Course {
	return &Course{l}
}

var a string

// "Getcourse is a"

func (c *Course) Getcourse(ctx context.Context, rr *protos.Request) (*protos.Response, error) {

	a := rr.GetKey()

	b, _, _ := time.Now().Clock()

	date = time.Now().Format("01-02-2006")
	count = 0

	db, err := sql.Open("mysql", "root:Akash@8888@tcp(127.0.0.1:3306)/course")
	fmt.Println("error")
	if err != nil {
		panic(err.Error())
		fmt.Println("err")
	}
	defer db.Close()

	db.QueryRow("SELECT name from courses where name = ? ", a).Scan(&result)
	db.QueryRow("SELECT count from details where name = ? and hour = ? and date=? ", a, b, date).Scan(&count)

	if result == a {
		switch {
		case (count == 0):
			c := 1
			_, err = db.Query("INSERT INTO details (name,value,count,hour,date) values(?,'available',?,?,?)", a, c, b, date)
			fmt.Println("insert", a, c, b, date)
		case (count != 0):
			_, err = db.Query("UPDATE details SET count = count+1 where hour= ? AND name= ? AND date=? ", b, a, date)
			fmt.Println("updated", count)

		}

	} else {
		fmt.Println("course not updated")
	}

	err = db.QueryRow("SELECT * from details where name = ? ORDER BY id DESC", a).Scan(&id, &name, &value, &count, &hour, &date)
	if err != nil {
		panic(err.Error())

	}

	return &protos.Response{Value: value, Count: count, Hour: hour}, nil
}
