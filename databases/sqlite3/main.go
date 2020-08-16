package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./foo.db")
	checkErr(err)

	//Insert
	smtm, err := db.Prepare("INSERT INTO userinfo(username, department, created) values(?,?,?)")
	checkErr(err)

	res, err := smtm.Exec("Neil", "Informatica", "2020-08-15")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
	//Update
	smtm, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	//Consult
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)
	var uid int
	var username string
	var department string
	var created time.Time

	for rows.Next() {
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	rows.Close()

	//Delete
	smtm, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err = smtm.Exec(id)
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

/* TRANSACTIONS */
/*  trashSQL, err := database.Prepare("update task set is_deleted='Y',last_modified_at=datetime() where id=?")
    if err != nil {
        fmt.Println(err)
    }
    tx, err := database.Begin()
    if err != nil {
    	fmt.Println(err)
    }
    _, err = tx.Stmt(trashSQL).Exec(id)
    if err != nil {
    	fmt.Println("doing rollback")
    	tx.Rollback()
    } else {
    	tx.Commit()
    }*/
