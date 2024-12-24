package api

import (
	"net/http"
	"strconv"

	"waifu-api/db"
	"waifu-api/models"

	"github.com/labstack/echo/v4"
)

// GetAllWaifus
func GetAllWaifus(c echo.Context) error {
    // Mengambil koneksi database
    db, err := db.ConnectDB()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal mengambil koneksi database"})
    }
    defer db.Close()

    // Mengambil daftar semua waifu dari database
    waifus, err := models.GetAllWaifus(db)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal mengambil daftar waifu"})
    }

    return c.JSON(http.StatusOK, waifus)
}

// GetWaifuByID mengembalikan waifu berdasarkan ID
func GetWaifuByID(c echo.Context) error {
    // Mengambil koneksi database
    db, err := db.ConnectDB()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal mengambil koneksi database"})
    }
    defer db.Close()

    // Mengambil ID waifu dari parameter URL
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "ID waifu tidak valid"})
    }

    // Mengambil waifu dari database berdasarkan ID
    waifu, err := models.GetWaifuByID(db, id)
    if err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"message": "Waifu tidak ditemukan"})
    }

    return c.JSON(http.StatusOK, waifu)
}

// CreateWaifu 
func CreateWaifu(c echo.Context) error {
    // Mengambil koneksi database
    db, err := db.ConnectDB()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal mengambil koneksi database"})
    }
    defer db.Close()

    // Binding data waifu dari request body
    var waifu models.Waifu
    if err := c.Bind(&waifu); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "Data waifu tidak valid"})
    }

    // Menambahkan waifu baru ke database
    err = models.CreateWaifu(db, &waifu)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal membuat waifu"})
    }

    return c.JSON(http.StatusCreated, waifu)
}

// UpdateWaifu 
func UpdateWaifu(c echo.Context) error {
    db, err := db.ConnectDB()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal mengambil koneksi database"})
    }
    defer db.Close()

    // Mengambil ID waifu dari parameter URL
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "ID waifu tidak valid"})
    }

    // Binding data waifu dari request body
    var waifu models.Waifu
    if err := c.Bind(&waifu); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "Data waifu tidak valid"})
    }

    // Memperbarui informasi waifu dalam database
    err = models.UpdateWaifu(db, id, &waifu)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal memperbarui waifu"})
    }

    return c.JSON(http.StatusOK, waifu)
}

// DeleteWaifu 
func DeleteWaifu(c echo.Context) error {
    // Mengambil koneksi database
    db, err := db.ConnectDB()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal mengambil koneksi database"})
    }
    defer db.Close()

    // Mengambil ID waifu dari parameter URL
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "ID waifu tidak valid"})
    }

    // Menghapus waifu dari database
    err = models.DeleteWaifu(db, id)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal menghapus waifu"})
    }

    return c.JSON(http.StatusOK, map[string]string{"message": "Waifu berhasil dihapus"})
}