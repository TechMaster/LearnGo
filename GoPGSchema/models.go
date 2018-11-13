package main

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"fmt"
)

type Article struct {
	tableName struct{} //`sql:"articles"`
	Id     int // Id is automatically detected as primary key
	Name   string
	Author   *Author
}

type Author struct {
	tableName struct{} //`sql:"person"`
	ID    int     // both "Id" and "ID" are detected as primary key
	Name  string  `sql:",unique"`
	Article []*Article // has many relation

}

var Db *pg.DB;
//TODO: get info from config
func ConnectDb() error {
	Db = pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "123",
		Database: "demo",
		Addr:     "localhost:5432",
	})


	var n int
	_, err := Db.QueryOne(pg.Scan(&n), "SELECT 1")
	return err
}

func InitSchema() {
	var article Article
	var author Author

	// Tạo bảng
	for _, model := range []interface{}{&author, &article} {
		err := Db.CreateTable(model, &orm.CreateTableOptions{
			Temp:          false,
			FKConstraints: true,
			IfNotExists:   true,
		})
		if err != nil {
			panic(err)
		}
	}
}

func InserData() {
	auth1 := Author {
		Name:"Cuong",
	}
	auth2 := Author {
		Name:"Dzung",
	}

	Db.Insert(&auth1)
	Db.Insert(&auth2)

	res, _ := orm.NewQuery(Db, &auth1).Insert()

	fmt.Println(auth2.ID)
	fmt.Println(res)
	/*if res, err := Db.Model(auth1, auth2).Insert(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}*/

	Db.Insert()

	a1 := Article{
		Name: "Gone with the Wind",
		Author:&auth1,
	}

	a2 := Article{
		Name: "Secret weapon in USA",
		Author:&auth2,
	}

	a3 := Article{
		Name: "Ganster in Newyork",
		Author:&auth1,
	}
	Db.Model(a1, a2, a3).Insert()
}