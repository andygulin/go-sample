package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

/**
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(20) DEFAULT NULL,
  `age` int(11) DEFAULT NULL,
  `address` varchar(100) DEFAULT NULL,
  `createdAt` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
*/

func main() {
	db, _ := sql.Open("mysql", "root:root@tcp(localhost:3306)/test?charset=utf8")
	defer db.Close()

	stmt, _ := db.Prepare("INSERT INTO user VALUES(NULL,?,?,?,?)")
	stmt.Exec("呵呵", 11, "上海", time.Now())
	stmt.Close()

	row := db.QueryRow("SELECT * FROM user ORDER BY id DESC LIMIT 1")
	var id, age int
	var name, address string
	var createdAt time.Time
	row.Scan(&id, &name, &age, &address, &createdAt)
	fmt.Println(id, name, age, address, createdAt)
}
