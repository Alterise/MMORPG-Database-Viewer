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
		msg, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			return
		}

		msg = strings.Trim(msg, "\b\r\n")

		args := strings.Split(msg, " ")
		cmd := strings.TrimSpace(args[0])
		args = args[1:]

		switch cmd {
		case "update_locations":
			c.commands <- command{
				id:     CMD_UPDATE_VIEW_LOCATIONS,
				client: c,
			}
		case "update_players":
			c.commands <- command{
				id:     CMD_UPDATE_VIEW_PLAYERS,
				client: c,
			}
		case "update_accounts":
			c.commands <- command{
				id:     CMD_UPDATE_VIEW_ACCOUNTS,
				client: c,
				args: args,
			}
		case "delete_account":
			c.commands <- command{
				id:     CMD_DELETE_ACCOUNT,
				client: c,
				args: args,
			}
		case "reset_password":
			c.commands <- command{
				id:     CMD_RESET_PASSWORD,
				client: c,
				args: args,
			}
		case "query":
			c.commands <- command{
				id:     CMD_QUERY,
				client: c,
				args: args,
			}
		case "verify":
			c.commands <- command{
				id:     CMD_VERIFY,
				client: c,
				args: args,
			}
		case "get_characters":
			c.commands <- command{
				id:     CMD_GET_CHARACTER_INFO,
				client: c,
				args: args,
			}
		case "quit":
			c.commands <- command{
				id:     CMD_QUIT,
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