package main

import (
	"bufio"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"net"
	"reflect"
)

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
		func() (int, int) { return len(data), len(names) },
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

	return container.NewHScroll(
		container.NewVScroll(
			container.NewGridWithRows(
				len(data) + 2,
				rows...
			),
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