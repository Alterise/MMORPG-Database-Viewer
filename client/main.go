package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	_ "github.com/lib/pq"
	"net"
	"strings"
)

var window fyne.Window
var conn net.Conn
var entryString binding.String
var entryStringLogin binding.String
var entryStringPassword binding.String

func updateLocationsView() {
	serverMsg := sendCommandToServer("update_locations")
	array := parseJsonToLocationInfoArray(serverMsg)
	window.SetContent(container.NewGridWithRows(2, NewLocationInfoView(array), createControlsView()))
}

func updatePlayersView() {
	serverMsg := sendCommandToServer("update_players")
	array := parseJsonToPlayerInfoArray(serverMsg)
	window.SetContent(container.NewGridWithRows(2, NewPlayerInfoView(array), createControlsView()))
}

func updateEmptyView() {
	window.SetContent(
		container.NewGridWithRows(
			2,
			widget.NewLabelWithStyle("Nothing in here yet", fyne.TextAlignCenter, fyne.TextStyle{}),
			createControlsView()),
	)
}

func updateErrorView(error string) {
	window.SetContent(
			widget.NewLabelWithStyle(error, fyne.TextAlignCenter, fyne.TextStyle{}),
	)
}

func updatePlayerView(login string) {
	serverMsg := sendCommandToServer("get_characters " + login)
	array := parseJsonToCharacterInfoArray(serverMsg)
	window.SetContent(
		container.NewGridWithRows(
			2,
			NewPlayerProfileView(array),
			widget.NewButton("Quit", func() {
				updateAuthView()
			}),
		),
	)
}

func updateAuthView() {
	entryStringLogin = binding.NewString()
	entryStringPassword = binding.NewString()
	window.SetContent(
		container.NewGridWithRows(
			2,
			widget.NewForm(
				widget.NewFormItem("Login:", widget.NewEntryWithData(entryStringLogin)),
				widget.NewFormItem("Password:", widget.NewEntryWithData(entryStringPassword)),
			),
			widget.NewButton("Login", func() {
				strLogin, _ := entryStringLogin.Get()
				strPassword, _ := entryStringPassword.Get()
				strLogin = strings.Trim(strLogin, " ")
				strPassword = strings.Trim(strPassword, " ")
				response := verifyUser(strLogin, strPassword)
				switch response {
				case AUTH_SUCCESS:
					updatePlayerView(strLogin)
				case AUTH_FAILED:
					dialog.ShowInformation("Error", "No such user.\nTry Again", window)
				case AUTH_ADMIN:
					updateEmptyView()
					dialog.ShowInformation("Greetings", "Welcome aboard, Captain", window)
				}
			}),
		),
	)
}

func main() {
	var err error
	conn, err = net.Dial("tcp", ":8888")
	if err == nil {
		defer conn.Close()
	}

	a := app.New()
	window = a.NewWindow("Character Viewer")
	window.Resize(fyne.NewSize(800, 600))

	if err != nil {
		updateErrorView("Cant connect to the server\n Try again Later")
	} else {
		updateAuthView()
	}

	window.ShowAndRun()

	if err == nil {
		sendCommandToServer("quit")
	}
}

