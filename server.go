package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
)

// Un servidor tiene varios rooms (salas)
type server struct {
	rooms    map[string]*room
	commands chan command
}

func (s *server) run() {
	for cmd := range s.commands {
		switch cmd.id {
		case CMD_NICK:
			s.nick(cmd.client, cmd.arg)
		case CMD_JOIN:
			s.join(cmd.client, cmd.arg)
		case CMD_ROOMS:
			s.listRooms(cmd.client, cmd.arg)
		case CMD_MSG:
			s.msg(cmd.client, cmd.arg)
		case CMD_QUIT:
			s.quit(cmd.client, cmd.arg)
		}
	}
}

func (s *server) nick(c *client, args []string) {
	c.nickname = args[1]
	c.msg(fmt.Sprintf("Bueno, te llamaré %s junior", c.nickname))
}

func (s *server) join(c *client, args []string) {
	roomName := args[1]

	// Chequea si el room existe. Si no, lo crea
	r, ok := s.rooms[roomName]
	if !ok {
		r = &room{
			name:    roomName,
			members: make(map[net.Addr]*client),
		}
		s.rooms[roomName] = r
	}

	r.members[c.connection.RemoteAddr()] = c

	s.quitCurrentRoom(c)
	// Asigna la sala actual al cliente
	c.currentRoom = r

	// Mensaje a la sala
	r.broadcast(c, fmt.Sprintf("%s se unió a la sala! \n", c.nickname))
	// Mensaje al cliente
	c.msg(fmt.Sprintf("Bienvenido a la sala %s ! Disfrute su estadía \n", r.name))
}

func (s *server) listRooms(c *client, args []string) {
	var rooms []string
	for name := range s.rooms {
		rooms = append(rooms, name)
	}
	c.msg(fmt.Sprintf("Salas disponibles : %s", strings.Join(rooms, ", ")))
}

func (s *server) msg(c *client, args []string) {
	if c.currentRoom == nil {
		c.err(errors.New("Tenés que entrar a una sala primero"))
		return
	}

	c.currentRoom.broadcast(c, c.nickname + ": "+strings.Join(args[1:len(args)], " "))
}

func (s *server) quit(c *client, args []string) {
	log.Printf("El cliente se ha desconectado: %s", c.connection.RemoteAddr().String())

	s.quitCurrentRoom(c)

	c.msg("Nos re vimos amigo ")
	c.connection.Close()
}

func (s *server) quitCurrentRoom(c *client) {
	if c.currentRoom != nil {
		delete(c.currentRoom.members, c.connection.RemoteAddr())
		c.currentRoom.broadcast(c, fmt.Sprintf("%s se fue de la sala =( ... ", c.nickname))
	}

}

func (s *server) newClient(conn net.Conn) {
	log.Printf("Nuevo cliente conectado : %s \n", conn.RemoteAddr().String())
	c := &client{
		connection: conn,
		nickname:   "anonymous",
		commands:   s.commands,
	}
	c.readInput()
}

// Creamos un Server
func newServer() *server {
	return &server{
		rooms: make(map[string]*room),
	}
}
