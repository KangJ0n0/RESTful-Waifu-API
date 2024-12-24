package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"

	"waifu-api/api"
	"waifu-api/db"
)

func main() {
    db, err := db.ConnectDB()
    if err != nil {
        log.Fatal("Gagal terkoneksi ke database:", err)
    }
    defer db.Close()
    e := echo.New()

    // Routing
    e.GET("/books", api.GetAllWaifus)
    e.GET("/books/:id", api.GetWaifuByID)
    e.POST("/books", api.CreateWaifu)
    e.PUT("/books/:id", api.UpdateWaifu)
    e.DELETE("/books/:id", api.DeleteWaifu)

    // Menjalankan server
    e.Logger.Fatal(e.Start(":8080"))
}