package main

type commandID int

const (
	CMD_UPDATE commandID = iota
	CMD_QUIT
)

type command struct {
	id commandID
	client *client
	args [] string
}