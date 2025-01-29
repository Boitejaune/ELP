package tcpclient

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

// Connection au serveur
func connectclient() {
    conn, err := net.Dial("tcp", "127.0.0.1:8080") 
    if err != nil {
        log.Fatal(err)
        os.Exit(1)
    }
    defer conn.Close()

    entree:=bufio.NewReader(os.Stdin)
    for {
        fmt.Print(">> ")
        text, err1:= entree.ReadString('\n') //Lit input client
        if err1 != nil {
            log.Println(err1)
            os.Exit(1)
        }

        fmt.Fprintf(conn, text+"\n") // Fprintf : imprime text dans c

        serveur,err2:= bufio.NewReader(conn).ReadString('\n') //Lit réponse serveur
        if err2 != nil {
            log.Println(err2)
            os.Exit(1)
        }

        fmt.Print("->: " + serveur) 


        if strings.TrimSpace(string(text)) == "STOP" {
                fmt.Println("TCP client exiting...")
                return
        }
    }

    conn.Close()
}



/*
//Connect
conn, err := net.Dial("tcp", port) //portString: “IP:Port”, eg “127.0.0.1:80”
// Disconnect
conn.Close()
defer conn.Close()
// Get yourself a reader on the connection, read some characters
reader := bufio.NewReader(conn)
message:= reader.ReadString(‘\n’)
// Write content on the connection
io.WriteString(conn, fmt.Sprintf("Coucou %d\n", i))

*/