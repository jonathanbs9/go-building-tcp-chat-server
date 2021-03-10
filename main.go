package main

import (
	"log"
	"net"
)

func main(){
	// Inicializamos el server
	s := newServer()
	go s.run()

	listener, err := net.Listen("tcp", ":8888")
	if err != nil{
		log.Fatalf("Error al inicializar el servidor : %s \n", err.Error())
	}
	defer listener.Close()
	log.Printf("Server inicializado en puerto :8888 ")

	for {
		conn, err := listener.Accept()
		if err != nil{
			log.Printf("No puede aceptar conexiones: %s \n", err.Error())
			continue
		}
		// Le asignamos al server la connection. (Go Routine)
		go s.newClient(conn)
	}
}
