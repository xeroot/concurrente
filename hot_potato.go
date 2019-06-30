package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func clearConsole() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

var servers = map[string]string{
	"Pedro": "8001",
	"Juan":  "8002",
	"Raul":  "8003",
	"Jorge": "8004"}

func startServer(nombre string, serverStarted chan bool) {
	ln, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", servers[nombre]))
	checkError(err)
	defer ln.Close()
	serverStarted <- true
	turno := make(chan bool, 1)
	for {
		con, err := ln.Accept()
		checkError(err)
		go handle(con, turno)
	}
}

func handle(con net.Conn, turno chan bool) {
	turno <- true
	r := bufio.NewReader(con)
	msg, err := r.ReadBytes('\n')
	checkError(err)
	msgStr := strings.TrimSpace(string(msg))
	if !checkSiAlguienPerdio(msgStr) {
		num, err := strconv.Atoi(msgStr)
		checkError(err)
		fmt.Println("Num recibido: ", num)
		if num == 0 {
			fmt.Println("PERDI :'(")
			avisarQuePerdi()
		} else {
			enviarNumeroAJugadorAleatorio(num - 1)
		}
	} else {
		fmt.Println("El jugador '", msgStr, "' perdio. :V")
		os.Exit(0)
	}
	<-turno
}

func checkSiAlguienPerdio(msgStr string) bool {
	for key := range servers {
		if key == msgStr {
			return true
		}
	}
	return false
}

func avisarQuePerdi() {
	for jugador, port := range servers {
		if jugador != nombre {
			go func(port string) {
				con, err := net.Dial("tcp", fmt.Sprintf("localhost:%s", port))
				checkError(err)
				defer con.Close()
				fmt.Fprintln(con, nombre)
			}(port)
		}
	}
}

func escogerRandomServer() string {
	randomN := rand.Intn(len(servers))
	i := 0
	for _, port := range servers {
		if i == randomN {
			return port
		}
		i++
	}
	return ""
}

func enviarNumeroAJugadorAleatorio(num int) {
	siguienteJugador := escogerRandomServer()
	for siguienteJugador == servers[nombre] {
		siguienteJugador = escogerRandomServer()
	}
	con, err := net.Dial("tcp", fmt.Sprintf("localhost:%s", siguienteJugador))
	checkError(err)
	defer con.Close()
	time.Sleep(time.Second * 1)
	fmt.Fprintln(con, num)
}

var nombre string

func main() {
	exit := make(chan bool)
	clearConsole()
	rand.Seed(time.Now().UTC().UnixNano())
	for nombre := range servers {
		fmt.Println(nombre)
	}
	fmt.Print("Escoja un personaje: ")
	fmt.Scan(&nombre)
	serverStarted := make(chan bool)
	go startServer(nombre, serverStarted)
	<-serverStarted
	fmt.Println("Servidor Iniciado.")
	fmt.Print("Empezar?(y/n): ")
	var empezar string
	fmt.Scan(&empezar)
	if empezar == "y" {
		fmt.Print("Escoja un numero: ")
		var num int
		fmt.Scan(&num)
		enviarNumeroAJugadorAleatorio(num - 1)
	}
	<-exit
}
