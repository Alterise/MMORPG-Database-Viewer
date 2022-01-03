package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"net"
)

var db *sql.DB

func main() {
	connectionString := fmt.Sprintf("host = %s port = %s user = %s password = %s dbname = %s sslmode = disable",
		host, port, user, password, dbname)

	var err error
	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()


	s := newServer()
	go s.run()

	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		go s.newClient(conn)
	}
	//fmt.Println(string(createLocationInfoJsonArray(db)))
}

