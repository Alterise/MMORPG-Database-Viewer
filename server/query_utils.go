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

func createLocationInfoJsonArray(db *sql.DB, args []string) []byte {
	sqlRows, err := db.Query("select * from locations order by location_id")
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

func createPlayerInfoJsonArray(db *sql.DB, args []string) []byte {
	sqlRows, err := db.Query(
		"select personal_data.username, personal_data.date_of_birth, players.\"level\", players.nickname, " +
			"players.\"race\" from personal_data, players, character_ownership where " +
			"character_ownership.account_id = personal_data.account_id and character_ownership.player_id = players.player_id")
	if err != nil {
		fmt.Println(err.Error())
	}

	var rows []PlayerInfo
	for sqlRows.Next() {
		var result PlayerInfo
		err = sqlRows.Scan(&result.Username, &result.DateOfBirth, &result.Level, &result.Nickname, &result.Race)
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

func createCharacterInfoJsonArray(db *sql.DB, args []string) []byte {
	login := args[0]
	characterInfoResult, err := db.Query(
		"select players.nickname, players.\"race\", players.\"level\", unique_items.item_name FROM unique_items, " +
			"personal_data, players, character_ownership where character_ownership.account_id = personal_data.account_id " +
			"and character_ownership.player_id = players.player_id and players.player_id = unique_items.owner_id and " +
			"personal_data.username = '" + login + "'")
	if err != nil {
		fmt.Println(err.Error())
	}

	var rows []CharacterInfo
	for characterInfoResult.Next() {
		var result CharacterInfo
		err = characterInfoResult.Scan(&result.Nickname, &result.Race, &result.Level, &result.ItemName)
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

