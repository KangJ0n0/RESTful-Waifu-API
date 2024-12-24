package models

import (
	"database/sql"
)

type Waifu struct {
    ID     int     `json:"id"`
    Name   string  `json:"name"`
    Age    int     `json:"age"`
    Height float64 `json:"height"`
    Weight float64 `json:"weight"`
    Series string  `json:"series"`
    Pict   string  `json:"pict"`
}

func GetAllWaifus(db *sql.DB) ([]Waifu, error) {
    rows, err := db.Query("SELECT id, name, age, height, weight, series, pict FROM waifus")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var waifus []Waifu
    for rows.Next() {
        var waifu Waifu
        if err := rows.Scan(&waifu.ID, &waifu.Name, &waifu.Age, &waifu.Height, &waifu.Weight, &waifu.Series, &waifu.Pict); err != nil {
            return nil, err
        }
        waifus = append(waifus, waifu)
    }

    return waifus, nil
}


func GetWaifuByID(db *sql.DB, id int) (Waifu, error) {
    var waifu Waifu
    err := db.QueryRow("SELECT id, name, age, height, weight, series, pict FROM waifus WHERE id = ?", id).Scan(&waifu.ID, &waifu.Name, &waifu.Age, &waifu.Height, &waifu.Weight, &waifu.Series, &waifu.Pict)
    if err != nil {
        return waifu, err
    }
    return waifu, nil
}


func CreateWaifu(db *sql.DB, waifu *Waifu) error {
    _, err := db.Exec("INSERT INTO waifus (name, age, height, weight, series, pict) VALUES (?, ?, ?, ?, ?, ?)", waifu.Name, waifu.Age, waifu.Height, waifu.Weight, waifu.Series, waifu.Pict)
    return err
}


func UpdateWaifu(db *sql.DB, id int, waifu *Waifu) error {
    _, err := db.Exec("UPDATE waifus SET name = ?, age = ?, height = ?, weight = ?, series = ?, pict = ? WHERE id = ?", waifu.Name, waifu.Age, waifu.Height, waifu.Weight, waifu.Series, waifu.Pict, id)
    return err
}


func DeleteWaifu(db *sql.DB, id int) error {
    _, err := db.Exec("DELETE FROM waifus WHERE id = ?", id)
    return err
}