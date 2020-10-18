package main

import (
	"database/sql"
	"fmt"

	"github.com/mazrean/gopendb-generator-proto/gopendb"
	"github.com/mazrean/gopendb-generator-proto/gopendb/users"
)

func main() {
	_db, _ := sql.Open("mysql", "")
	db := gopendb.NewDB(_db)
	condition := users.NewCondition(users.ID.Eq(0))
	users, _ := db.Users().Where(condition).Select(users.ID).Do()
	fmt.Println(users)
}
