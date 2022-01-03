package main

import (
	"fmt"
	"net"
)

type server struct {
	commands chan command
}

func newServer() *server {
	return &server{
		commands: make(chan command),
	}
}

func (s *server) run() {
	for cmd := range s.commands {
		switch cmd.id {
		case CMD_UPDATE:
			s.sendJson(cmd.client)
			s.quit(cmd.client)
		}
	}
}

func (s *server) sendJson(c *client) {
	json :=string(createLocationInfoJsonArray(db))
	c.msg(json)
	fmt.Println("sent: " + json)
}

func (s *server) newClient(conn net.Conn) {
	//fmt.Printf("new client has joined: %s\n", conn.RemoteAddr().String())

	c := &client{
		conn:     conn,
		commands: s.commands,
	}

	c.readCommand()
}

func (s *server) quit(c *client) {
	//fmt.Printf("client has disconnected: %s\n", c.conn.RemoteAddr().String())
	c.conn.Close()
}