package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// ConnectDB berfungsi untuk membuat koneksi ke database MySQL
func ConnectDB() (*sql.DB, error) {
	// 
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/waifu-test")
	if err != nil {
		return nil, err
	}

	// Memeriksa koneksi ke database
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Berhasil terkoneksi ke database MySQL")
	return db, nil
}