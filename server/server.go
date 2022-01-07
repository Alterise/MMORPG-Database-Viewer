package main

import (
	"database/sql"
	"fmt"
	"net"
	"strings"
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
		case CMD_UPDATE_VIEW_LOCATIONS:
			s.sendJson(cmd.client, cmd.args, createLocationInfoJsonArray)
		case CMD_UPDATE_VIEW_PLAYERS:
			s.sendJson(cmd.client, cmd.args, createPlayerInfoJsonArray)
		case CMD_UPDATE_VIEW_ACCOUNTS:
			s.sendJson(cmd.client, cmd.args, createAccountInfoJsonArray)
		case CMD_RESET_PASSWORD:
			s.resetPassword(cmd.client, cmd.args)
		case CMD_DELETE_ACCOUNT:
			s.deleteAccount(cmd.client, cmd.args)
		case CMD_QUERY:
			s.processQuery(cmd.client, cmd.args)
		case CMD_VERIFY:
			s.verify(cmd.client, cmd.args)
		case CMD_GET_CHARACTER_INFO:
			s.sendJson(cmd.client, cmd.args, createCharacterInfoJsonArray)
		case CMD_QUIT:
			s.quit(cmd.client)
		}
	}
}

func (s *server) sendJson(c *client, args []string, createFunc func(*sql.DB, []string) []byte) {
	json :=string(createFunc(db, args))
	c.msg(json)
	fmt.Println("sent: " + json)
}

func (s *server) deleteAccount(c *client, args []string) {
	_, err := db.Exec("DELETE FROM personal_data WHERE username ='" + args[0] + "'")
	if err != nil {
		c.msg("failed")
		fmt.Println("Can't delete user with username: " + args[0])
	} else {
		c.msg("success")
		fmt.Println("User with username: " + args[0] + " has been deleted")
	}
}

func (s *server) resetPassword(c *client, args []string) {
	_, err := db.Exec("UPDATE personal_data SET password_hash = " +
		"'cdb59355f3ba293977fc0945fb85f11822d412c45c7520c7121bd2234f6c1f48' WHERE username = '" + args[0] + "'")
	if err != nil {
		c.msg("failed")
		fmt.Println("Can't reset password for user with username: " + args[0])
	} else {
		c.msg("success")
		fmt.Println("Password for user with username: " + args[0] + " has been reset")
	}
}

func (s *server) processQuery(c *client, args []string) {
	query := strings.Join(args, " ")
	_, err := db.Exec(query)
	if err != nil {
		c.msg("failed")
		fmt.Println("Query execution failed: " + query)
	} else {
		c.msg("success")
		fmt.Println("Query executed successfully: " + query)
	}
}



func (s *server) verify(c *client, args []string) {
	login, password := args[0], args[1]
	passwordHashResult, err := db.Query("SELECT password_hash FROM personal_data WHERE username = '" + login + "'")
	var passwordHash string
	correctLogin := false
	for passwordHashResult.Next() {
		err = passwordHashResult.Scan(&passwordHash)
		correctLogin = true
	}
	correctPassword := passwordHash == password

	if err != nil {
		c.msg("query_failed")
		fmt.Println("Query failed")
	} else if correctLogin && correctPassword {
		c.msg("success")
		fmt.Println("User " + login + " with password hash " + password + " exists")
	} else {
		c.msg("failed")
		fmt.Println("User " + login + " with password hash " + password + " doesn't exist")
	}
}

func (s *server) newClient(conn net.Conn) {
	fmt.Printf("new client has joined: %s\n", conn.RemoteAddr().String())

	c := &client{
		conn:     conn,
		commands: s.commands,
	}

	c.readCommand()
}

func (s *server) quit(c *client) {
	fmt.Printf("client has disconnected: %s\n", c.conn.RemoteAddr().String())
	c.conn.Close()
}