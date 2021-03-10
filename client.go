package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

type client struct {
	connection  net.Conn
	nickname    string
	currentRoom *room
	commands    chan<- command
}

func (c *client) readInput() {
	for {
		msg, err := bufio.NewReader(c.connection).ReadString('\n')
		if err != nil {
			log.Fatal("Conexión perdida : %s", err.Error())
			return
		}

		msg = strings.Trim(msg, "\r\n")
		args := strings.Split(msg, " ")
		cmd := strings.TrimSpace(args[0])
		switch cmd {
		case "/nickname":
			c.commands <- command{
				id:     CMD_NICK,
				client: c,
				arg:    args,
			}
		case "/join":
			c.commands <- command{
				id:     CMD_JOIN,
				client: c,
				arg:    args,
			}
		case "/rooms":
			c.commands <- command{
				id:     CMD_ROOMS,
				client: c,
				arg:    args,
			}
		case "/msg":
			c.commands <- command{
				id:     CMD_MSG,
				client: c,
				arg:    args,
			}
		case "/quit":
			c.commands <- command{
				id:     CMD_QUIT,
				client: c,
				arg:    args,
			}
		default:
			c.err(fmt.Errorf("Comando desconocido : %s", cmd))
		}
	}
}

func (c *client) err(err error) {
	c.connection.Write([]byte("ERR: " + err.Error() + " \n"))
}

func (c *client) msg(msg string) {
	c.connection.Write([]byte("> " + msg + " \n"))
}
