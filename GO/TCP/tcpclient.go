packkge main

import (
    "bufio"
    "fmt"
    "net"
    "os"
    "strings"
)


func main() {
    conn, err := net.Dial("tcp", port) 
    if err != nil {
        log.Fatal(err)
    }
    defer l.Close()
}


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