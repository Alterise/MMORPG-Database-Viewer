package main

import (
	"encoding/json"
	"fyne.io/fyne/v2"
	"strconv"
)

type LocationInfo struct {
	Id    		int		`json:"id"`
	Name 		string  `json:"name"`
	Leader 		string 	`json:"leader"`
	Hostility 	int		`json:"hostility"`
}

func NewLocationInfoView(rows []LocationInfo) fyne.CanvasObject {
	var data [][]string
	for _, location := range rows {
		row := []string{
			strconv.Itoa(location.Id),
			location.Name,
			location.Leader,
			strconv.Itoa(location.Hostility),
		}
		data = append(data, row)
	}
	//return createTableView(LocationInfo{}, data)
	return createGridView(LocationInfo{}, data)
}

func parseJsonToLocationInfoArray(jsonString string) []LocationInfo {
	var array []LocationInfo
	err := json.Unmarshal([]byte(jsonString), &array)
	if err != nil {
		println(err.Error())
		return nil
	}

	return array
}