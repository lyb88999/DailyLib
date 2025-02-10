package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

func GetUserById(db *sqlx.DB, id int) (*User, error) {
	var user User
	err := db.Get(&user, "SELECT id, name FROM user WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetAllUsers(db *sqlx.DB) ([]*User, error) {
	var users []*User
	err := db.Select(&users, "SELECT id, name FROM user")
	if err != nil {
		return nil, err
	}
	return users, nil
}

func CreateUser(db *sqlx.DB, name string) error {
	tx := db.MustBegin()
	_, err := tx.Exec("INSERT INTO user (name) VALUES (?)", name)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func BulkInsertUsers(db *sqlx.DB, user []*User) error {
	query := "INSERT INTO user (name) VALUES (:name)"
	_, err := db.NamedExec(query, user)
	return err
}
func main() {
	var err error
	db, err := sqlx.Open("mysql", "root:Lyb1217..@tcp(127.0.0.1:3306)/badminton")
	if err != nil {
		log.Fatalln(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Connected to the database")
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(db)
	if err = CreateUser(db, "John Doe"); err != nil {
		log.Fatalln(err)
	}
	if err = BulkInsertUsers(db, []*User{
		{Name: "Alice"},
		{Name: "Bob"},
		{Name: "Charlie"},
	}); err != nil {
		log.Fatalln(err)
	}
	if users, err := GetAllUsers(db); err != nil {
		log.Fatalln(err)
	} else {
		for _, user := range users {
			fmt.Printf("User ID: %d, Name: %s\n", user.ID, user.Name)
		}
	}
	if user, err := GetUserById(db, 1); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("User ID: %d, Name: %s\n", user.ID, user.Name)
	}

}
