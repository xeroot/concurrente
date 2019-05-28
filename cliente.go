package main
import (
	"fmt"
	"net"
	"bufio"
	"os"
)

func main() {
	fmt.Println("Iniciando Cliente...")
	con,_ := net.Dial("tcp","localhost:8000")
	defer con.Close()
	r := bufio.NewReader(con)
	gin := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Mensaje: ")
		msg,_ := gin.ReadString('\n')
		fmt.Fprint(con,msg)
		resp,_ := r.ReadString('\n')
		fmt.Print("Respuesta: ",resp)
		if msg[0] == 'x' {
			break;
		}
	}
}