package main

import (
	"fmt"
	"sync"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 3000; i++ {
		wg.Add(1)
		go insert(i, wg)
	}
	wg.Wait()
}

func insert(i int, wg *sync.WaitGroup) {
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/example")
	if err != nil {
		fmt.Println("failed to open database connection")
		panic(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		fmt.Println("failed to begin transaction")
		panic(err)
	}
	_, err = db.Exec("INSERT INTO test(uniquekey) values (?)", "same")
	if err != nil {
		fmt.Println("failed to exec show tables")
		panic(err)
	}
	if err := tx.Rollback(); err != nil {
		fmt.Println("failed to rollback")
		panic(err)
	}

	wg.Done()
}
