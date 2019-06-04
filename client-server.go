package main
import (
	"fmt"
	"net"
	"bufio"
	"time"
)

func handle(con net.Conn) {
	defer con.Close()
	r := bufio.NewReader(con)
	for {
		msg, _ := r.ReadString('\n')
		fmt.Print("Recibido: ",msg)
		SendMessageTo(msg)
	}
}

func StartServer() {
	fmt.Println("Iniciando Servidor...")
	ln, _:= net.Listen("tcp", "10.11.97.219:9000")
	defer ln.Close()
	fmt.Println("Escuchando por puerto 9000")
	for {
		con, _ := ln.Accept()
		go handle(con)
	}
}

func SendMessageTo(msg string) {
	con,_ := net.Dial("tcp","10.11.97.218:9000")
	defer con.Close()
	fmt.Fprint(con,msg)
}

func main() {
	//StartServer()
	SendMessageTo("hola men")
	time.Sleep(time.Second)
}