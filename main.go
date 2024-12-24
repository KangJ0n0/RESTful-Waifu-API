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
    e.GET("/waifus", api.GetAllWaifus)
    e.GET("/waifus/:id", api.GetWaifuByID)
    e.POST("/waifus", api.CreateWaifu)
    e.PUT("/waifus/:id", api.UpdateWaifu)
    e.DELETE("/waifus/:id", api.DeleteWaifu)

    // Menjalankan server
    e.Logger.Fatal(e.Start(":8080"))
}