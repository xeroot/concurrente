package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
)

type Node struct {
	ConnectionPort string `json: "ConnectionPort"`
	DataPort       string `json: "DataPort"`
}

type Request struct {
	JustAdd bool `json: "JustAdd"`
	Node    Node `json: "Node"`
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func startConexionServer(serverIsReady chan bool, connectionPort string) {
	ln, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", connectionPort))
	checkError(err)
	defer ln.Close()
	serverIsReady <- true
	freeWorking := make(chan bool, 1)
	freeWorking <- true
	for {
		con, err := ln.Accept()
		checkError(err)
		go conexionHandle(con, freeWorking)
	}
}

func conexionHandle(con net.Conn, freeWorking chan bool) {
	defer con.Close()
	<-freeWorking
	r := bufio.NewReader(con)
	resp, err := r.ReadBytes('\n')
	checkError(err)
	var req Request
	json.Unmarshal(resp, &req)
	newNode := req.Node
	if !req.JustAdd {
		allNodes := append(nodes, thisNode)
		allJSONNodes, err := json.Marshal(allNodes)
		checkError(err)
		fmt.Fprintln(con, string(allJSONNodes)) // enviar los nodos que se, al nuevo
		// enviar el nuevo, a los demas nodos
		newRequest := Request{JustAdd: true, Node: newNode}
		for _, node := range nodes {
			go sendNode(node.ConnectionPort, newRequest)
		}
		nodes = append(nodes, newNode)
	} else {
		nodes = append(nodes, newNode)
	}
	fmt.Println("NODES: ", nodes)
	freeWorking <- true
}

func sendNode(destinationPort string, newRequest Request) {
	con, err := net.Dial("tcp", fmt.Sprintf("localhost:%s", destinationPort))
	checkError(err)
	defer con.Close()
	newNode, err := json.Marshal(newRequest)
	checkError(err)
	fmt.Fprintln(con, string(newNode))
}

func startDataServer(serverIsReady chan bool, dataPort string) {
	ln, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", dataPort))
	checkError(err)
	defer ln.Close()
	serverIsReady <- true
	for {
		con, err := ln.Accept()
		checkError(err)
		go datahandle(con)
	}
}

func datahandle(con net.Conn) {
	defer con.Close()
}

func connectToNode(destConPort string) {
	con, err := net.Dial("tcp", fmt.Sprintf("localhost:%s", destConPort))
	checkError(err)
	defer con.Close()
	r := bufio.NewReader(con)
	request := Request{JustAdd: false, Node: thisNode}
	reqJSONObj, err := json.Marshal(request)
	checkError(err)
	//fmt.Println("Enviando...", string(reqJSONObj))
	fmt.Fprintln(con, string(reqJSONObj))
	resp, err := r.ReadBytes('\n')
	checkError(err)
	//fmt.Print("Respuesta: ", string(resp[:]))
	json.Unmarshal(resp, &nodes)
	fmt.Println("NODES: ", nodes)
}

var connectionPort string
var dataPort string
var nodes = make([]Node, 0)
var thisNode Node

func main() {
	exit := make(chan bool)
	fmt.Println("Iniciando Nodo...")
	fmt.Print("Seleccione Puerto de conexiones(800N): ")
	//var connectionPort string
	fmt.Scan(&connectionPort)
	fmt.Print("Seleccione Puerto de datos(900N): ")
	//var dataPort string
	fmt.Scan(&dataPort)
	serversChan := make(chan bool, 2)
	fmt.Println("Iniciando Servidores...")
	go startConexionServer(serversChan, connectionPort)
	go startDataServer(serversChan, dataPort)
	<-serversChan
	<-serversChan
	fmt.Println("Servidores Iniciados.")
	thisNode = Node{ConnectionPort: connectionPort, DataPort: dataPort}
	fmt.Print("Unirse a otro nodo? (y/n): ")
	var join string
	fmt.Scan(&join)
	if join == "y" {
		fmt.Print("Puerto a unirse(800N): ")
		var nextNodePort string
		fmt.Scan(&nextNodePort)
		connectToNode(nextNodePort)
	}
	// :v
	<-exit
}
