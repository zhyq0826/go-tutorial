package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	//"time"
)

// user@unix(/path/to/socket)/dbname?charset=utf8
// user:password@tcp(localhost:5555)/dbname?charset=utf8
// user:password@/dbname
// user:password@tcp([de:ad:be:ef::ca:fe]:80)/dbname

func main() {
	db, err := sql.Open("mysql", "root:@(localhost:3306)/sqlalchemy_lab?charset=utf8")
	checkErr(err)

	//插入数据
	stmt, err := db.Prepare("INSERT into tag (id, name, group_id) values(0, ?, ?)")
	checkErr(err)

	res, err := stmt.Exec("研发部门", 10)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
	//更新数据
	stmt, err = db.Prepare("update tag set name=? where id=?")
	checkErr(err)

	res, err = stmt.Exec("研发部门update", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	//查询数据
	rows, err := db.Query("SELECT * FROM tag order by id desc limit 3")
	checkErr(err)

	for rows.Next() {
		var id int
		var name string
		var groupId string
		err = rows.Scan(&id, &name, &groupId)
		checkErr(err)
		fmt.Println(id)
		fmt.Println(name)
		fmt.Println(groupId)
	}

	//删除数据
	stmt, err = db.Prepare("delete from tag where id=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
