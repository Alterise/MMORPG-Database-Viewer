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

type PlayerInfo struct {
	Username    string 	`json:"username"`
	DateOfBirth string  `json:"dateOfBirth"`
	Level 		int 	`json:"level"`
	Nickname 	string	`json:"nickname"`
	Race		string	`json:"race"`
}

type CharacterInfo struct {
	Nickname 	string	`json:"nickname"`
	Race		string	`json:"race"`
	Level 		int 	`json:"level"`
	ItemName    string 	`json:"itemName"`
}

type AccountInfo struct {
	Id    		int		`json:"id"`
	Username    string 	`json:"username"`
	DateOfBirth string  `json:"dateOfBirth"`
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
	return createTableView(LocationInfo{}, data)
	//return createGridView(LocationInfo{}, data)
}

func NewPlayerInfoView(rows []PlayerInfo) fyne.CanvasObject {
	var data [][]string
	for _, player := range rows {
		row := []string{
			player.Username,
			player.DateOfBirth,
			strconv.Itoa(player.Level),
			player.Nickname,
			player.Race,
		}
		data = append(data, row)
	}
	return createTableView(PlayerInfo{}, data)
	//return createGridView(PlayerInfo{}, data)
}

func NewAccountInfoView(rows []AccountInfo) fyne.CanvasObject {
	var data [][]string
	for _, account := range rows {
		row := []string{
			strconv.Itoa(account.Id),
			account.Username,
			account.DateOfBirth,
		}
		data = append(data, row)
	}
	return createTableView(AccountInfo{}, data)
	//return createGridView(PlayerInfo{}, data)
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

func parseJsonToPlayerInfoArray(jsonString string) []PlayerInfo {
	var array []PlayerInfo
	err := json.Unmarshal([]byte(jsonString), &array)
	if err != nil {
		println(err.Error())
		return nil
	}

	return array
}

func parseJsonToAccountInfoArray(jsonString string) []AccountInfo {
	var array []AccountInfo
	err := json.Unmarshal([]byte(jsonString), &array)
	if err != nil {
		println(err.Error())
		return nil
	}

	return array
}

func parseJsonToCharacterInfoArray(jsonString string) []CharacterInfo {
	var array []CharacterInfo
	err := json.Unmarshal([]byte(jsonString), &array)
	if err != nil {
		println(err.Error())
		return nil
	}

	return array
}