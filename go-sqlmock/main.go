package main

import (
	"fmt"
	"log"

	"github.com/DATA-DOG/go-sqlmock"
)

func main() {
	// 创建mock数据库连接
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	// 设置期望的数据库查询
	mock.ExpectQuery("SELECT name FROM users WHERE id = ?").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("John Doe"))

	// 执行查询
	rows, err := db.Query("SELECT name FROM users WHERE id = ?", 1)
	if err != nil {
		log.Fatalf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var name string
	for rows.Next() {
		if err = rows.Scan(&name); err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}
		fmt.Println(name)
	}

	// 验证所有的期望是否都得到了满足
	if err = mock.ExpectationsWereMet(); err != nil {
		log.Fatalf("there were unfulfilled expectations: %v", err)
	}
}
