package Sql

import (
	"database/sql"
	"fmt"

	//"strconv"

	_ "github.com/go-sql-driver/mysql"
)

// type Taxis struct {
// 	//id     int    `json:"Id"`
// 	name   string `json:"Name"`
// 	number string `json:"Number"`
// 	price  uint   `json:"Price"`
// }

func AdedSql(a, b string, c int) { //Установка данных в MySql
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/taxi")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	//"insert into productdb.Products (model, company, price) values (?, ?, ?)",
	result, err := db.Exec("INSERT INTO taxi (name, number, price) VALUES(?,?,?)", a, b, c)
	if err != nil {
		panic(err)
	}

	fmt.Println(result.LastInsertId())
}
