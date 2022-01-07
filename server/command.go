package main

type commandID int

const (
	CMD_UPDATE_VIEW_LOCATIONS commandID = iota
	CMD_UPDATE_VIEW_PLAYERS
	CMD_UPDATE_VIEW_ACCOUNTS
	CMD_DELETE_ACCOUNT
	CMD_RESET_PASSWORD
	CMD_QUERY
	CMD_VERIFY
	CMD_GET_CHARACTER_INFO
	CMD_QUIT
)

type command struct {
	id commandID
	client *client
	args [] string
}