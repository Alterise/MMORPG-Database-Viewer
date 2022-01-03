package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type client struct {
	conn net.Conn
	commands chan<- command
}

func (c *client) readCommand() {
	for {
		cmd, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			return
		}

		cmd = strings.Trim(cmd, "\b\r\n")

		switch cmd {
		case "update":
			c.commands <- command{
				id:     CMD_UPDATE,
				client: c,
			}
		default:
			c.err(fmt.Errorf("unknown command: %s", cmd))
			c.conn.Close()
		}
	}
}

func (c *client) err(err error) {
	c.conn.Write([]byte("err: " + err.Error() + "\n"))
}

func (c *client) msg(msg string) {
	c.conn.Write([]byte(msg + "\n"))
}