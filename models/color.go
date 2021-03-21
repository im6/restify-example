package models

import (
	"errors"
	"fmt"

	"github.com/im6/vp3/store"
)

// Color is definition
type Color struct {
	Id    int64  `json:"k"`
	Star  int64   `json:"s"`
	Color string  `json:"v"`
}

// Get color collection from DB
func GetColors(queryType string) ([]Color, error) {
	var orderByField string
	switch queryType {
		case "latest":
			orderByField = "id"
	  case "popular":
			orderByField = "star"
	}
	queryStr := fmt.Sprintf(
		"SELECT id, color, star FROM colorpk_color WHERE display = 0 ORDER BY %s desc",
		orderByField,
	)
	var colors []Color
	db := store.GetDbConnection()
	rows, err := db.Query(queryStr)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
	  var c Color
		err = rows.Scan(&c.Id, &c.Color, &c.Star)
		if err != nil {
			return nil, err
		}
		colors = append(colors, c)
  }
	return colors, nil
}

func GetOneColor(colorId string) (*Color, error) {
	db := store.GetDbConnection()
	rows, err := db.Query("SELECT id, color, star FROM colorpk_color WHERE id = ?", colorId)
	defer rows.Close()
  if rows.Next() {
	  var c Color
		err = rows.Scan(&c.Id, &c.Color, &c.Star)
		if err != nil {
			return nil, err
		}
		return &c, nil
  } else {
		return nil, errors.New("invalid color id")
	}
}

func IncrementColorStar(id int64) (error) {
	db := store.GetDbConnection()
	res, err := db.Exec("UPDATE colorpk_color SET star = star + 1 WHERE id = ?", id)
	if err != nil {
    return err
	}
	rowNum, err := res.RowsAffected()
	if err != nil || rowNum != 1  {
		return errors.New("invalid multiple update")
	}
	return nil
}
