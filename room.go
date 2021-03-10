package main

import "net"

type room struct {
	name string
	members map[net.Addr] *client
}

func (r room) broadcast(sender *client, msg string) {
	for addr, m := range r.members{
		if addr != sender.connection.RemoteAddr(){
			m.msg(msg)
		}
	}
}
