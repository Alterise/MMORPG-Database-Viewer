package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

type LocationInfo struct {
	Id    		int		`json:"id"`
	Name 		string  `json:"name"`
	Leader 		string 	`json:"leader"`
	Hostility 	int		`json:"hostility"`
}

func createLocationInfoJsonArray(db *sql.DB) []byte {
	sqlRows, err := db.Query("SELECT * FROM locations ORDER BY location_id")
	if err != nil {
		fmt.Println(err.Error())
	}

	var rows []LocationInfo
	for sqlRows.Next() {
		var result LocationInfo
		err = sqlRows.Scan(&result.Id, &result.Name, &result.Leader, &result.Hostility)
		if err != nil {
			fmt.Println(err.Error())
		}
		rows = append(rows, result)
	}

	marshal, err := json.Marshal(rows)
	if err != nil {
		fmt.Println(err.Error())
	}

	return marshal
}