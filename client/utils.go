package main

import (
	"bufio"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"net"
	"reflect"
	"strconv"
)

func createControlsView() fyne.CanvasObject {
	entryString = binding.NewString()
	controlsView := container.NewGridWithRows(
		2,
		container.NewAdaptiveGrid(
			2,
			widget.NewButton("Locations", updateLocationsView),
			widget.NewButton("Players", updatePlayersView),
		),
		widget.NewForm(
			widget.NewFormItem(
				"Execute query: ",
				widget.NewButton("Execute", func() {
					str, _ := entryString.Get()
					sendCommandToServer("query " + str)
				}),
			),
			widget.NewFormItem("Query: ", widget.NewEntryWithData(entryString)),
		),
	)
	return controlsView
}

func getNames(currentStruct interface{}) []string {
	values := reflect.TypeOf(currentStruct)
	names := make([]string, values.NumField())
	for i := 0; i < values.NumField(); i++ {
		names[i] = values.Field(i).Name
	}
	return names
}

func createNamesRow(currentStruct interface{}) fyne.CanvasObject {
	names := getNames(currentStruct)
	var labelArray []fyne.CanvasObject
	for i := 0; i < len(names); i++ {
		labelArray = append(labelArray, widget.NewLabelWithStyle(names[i], fyne.TextAlignCenter, fyne.TextStyle{}))
	}
	return container.NewGridWithColumns(len(names), labelArray...)
}

func createTableView(cellStruct interface{}, data [][]string) fyne.CanvasObject {
	if len(data) == 0 {
		return widget.NewLabel("Table is empty")
	}
	names := getNames(cellStruct)

	table := widget.NewTable(
		func() (int, int) { return len(data) + 1, len(names) },
		func() fyne.CanvasObject {
			return widget.NewLabelWithStyle("Error", fyne.TextAlignCenter, fyne.TextStyle{})
		},
		func(id widget.TableCellID, cell fyne.CanvasObject) {
			label := cell.(*widget.Label)
			switch id.Row {
			case 0:
				label.SetText(names[id.Col])
			default:
				label.SetText(data[id.Row - 1][id.Col])
			}
		})
	for i := 0; i < len(names); i++ {
		table.SetColumnWidth(i, 250)
	}
	return table
}

func createGridView(cellStruct interface{}, data [][]string) fyne.CanvasObject {
	if len(data) == 0 {
		return widget.NewLabel("Table is empty")
	}
	names := getNames(cellStruct)

	var rows []fyne.CanvasObject
	rows = append(rows,
		widget.NewLabelWithStyle("Database", fyne.TextAlignCenter, fyne.TextStyle{}))
	rows = append(rows, createNamesRow(cellStruct))
	for i := 0; i < len(data); i++ {
		var cols []fyne.CanvasObject
		for j := 0; j < len(names); j++ {
			cols = append(cols, widget.NewLabelWithStyle(data[i][j], fyne.TextAlignCenter, fyne.TextStyle{}))
		}
		row := container.NewGridWithColumns(len(names), cols...)
		rows = append(rows, row)
	}

	return container.NewScroll(
		container.NewGridWithRows(
			len(data) + 2,
			rows...
		),
	)
}

func sendCommandToServer(cmd string) string {
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		println(err.Error())
		return ""
	}

	defer conn.Close()

	_, err = fmt.Fprintf(conn, cmd + "\n")
	if err != nil {
		println(err.Error())
		return ""
	}

	var message string
	message, err = bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		println(err.Error())
		return ""
	}

	return message
}

func createCharacterList(characterMap map[string][]CharacterInfo) []fyne.CanvasObject {
	var list []fyne.CanvasObject

	for key, characters := range characterMap {
		var itemLabels []fyne.CanvasObject
		for _, character := range characters {
			itemLabels = append(itemLabels, widget.NewLabelWithStyle(character.ItemName, fyne.TextAlignCenter, fyne.TextStyle{}))
		}
		list = append(list, container.NewGridWithColumns(
				4,
				widget.NewLabelWithStyle(key, fyne.TextAlignCenter, fyne.TextStyle{}),
				widget.NewLabelWithStyle(characters[0].Race, fyne.TextAlignCenter, fyne.TextStyle{}),
				widget.NewLabelWithStyle(strconv.Itoa(characters[0].Level), fyne.TextAlignCenter, fyne.TextStyle{}),
				container.NewVScroll(container.NewGridWithRows(len(itemLabels), itemLabels...)),
			),
		)
	}
	return list
}

func NewPlayerProfileView(array []CharacterInfo) fyne.CanvasObject {
	characterMap := make(map[string][]CharacterInfo)
	for _, player := range array {
		nickname := player.Nickname
		characterMap[nickname] = append(characterMap[nickname], player)
	}

	var list []fyne.CanvasObject
	list = append(list, container.NewGridWithColumns(
		4,
		widget.NewLabelWithStyle("Nickname:", fyne.TextAlignCenter, fyne.TextStyle{}),
		widget.NewLabelWithStyle("Race:", fyne.TextAlignCenter, fyne.TextStyle{}),
		widget.NewLabelWithStyle("Level:", fyne.TextAlignCenter, fyne.TextStyle{}),
		widget.NewLabelWithStyle("Inventory:", fyne.TextAlignCenter, fyne.TextStyle{}),
	))
	list = append(list, createCharacterList(characterMap)...)

	return container.NewScroll(
		container.NewGridWithRows(
			len(list),
			list...,
		),
	)
}