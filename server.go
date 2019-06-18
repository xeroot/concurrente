package main

import (
	"bufio"
	"fmt"
	"net"
)

func handle(con net.Conn) {
	defer con.Close()
	r := bufio.NewReader(con)
	for {
		// test theia xd
		// aqui se chanta esperando otro request (otra linea)
		msg, _ := r.ReadString('\n')
		fmt.Fprintln(con, msg) // aqui le responde al cliente
		fmt.Println("Recibido: ", msg)
		if len(msg) > 0 && msg[0] == 'x' {
			break
		}
	}
}

func main() {
	fmt.Println("Iniciando Servidor...")
	ln, _ := net.Listen("tcp", "localhost:8000")
	defer ln.Close()
	fmt.Println("Escuchando por puerto 8000")
	for {
		// aqui se quede esperando a un cli
		// OJO, no cada request sino un cliente
		con, _ := ln.Accept()
		go handle(con)
	}
}
