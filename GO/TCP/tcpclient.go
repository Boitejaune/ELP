packkge main

import (
    "bufio"
    "fmt"
    "net"
    "os"
    "strings"
)

// Connection au serveur
func main() {
    conn, err := net.Dial("tcp", port) 
    if err != nil {
        log.Fatal(err) // comment?????
        os.Exit()
    }
    defer l.Close() //????????

    entree:=bufio.NewReader(os.Stdin)
    for {
        fmt.Print(">> ")
        text, err1:= entree.ReadString('\n') //Lit input client
        if err1 != nil {
            log.Println(err1)
            os.Exit(1)
        }

        fmt.Fprintf(c, text+"\n") // Fprintf : imprime text dans c

        serveur,err2:= bufio.NewReader(c).ReadString('\n') //Lit réponse serveur
        if err2 != nil {
            log.Println(err2)
            os.Exit(1)
        }

        fmt.Print("->: " + message) 


        if strings.TrimSpace(string(text)) == "STOP" {
                fmt.Println("TCP client exiting...")
                return
        }
    }

    c.Close()
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