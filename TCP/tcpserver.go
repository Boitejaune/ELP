package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
    "strings"
)

const (
    HOST = "localhost"
    port = ":8080"
    TYPE = "tcp"
)

// connection au serveur
func main() {
    l, err := net.Listen("tcp", port)  //Take a TCP port on the machine and ask connection attempts to that port to be redirected to your app
    if err != nil {
        log.Fatal(err)
    }
    defer l.Close()

    for {
        c, err := l.Accept()
        if err != nil {
            log.Fatal(err)
        }
        go connection(c)
    }
}


// Gestion connection
func connection( c net.Conn){

}



// Accept a new connection on that port
conn, errconn := ln.Accept()
// Close the connection of a client
conn.Close()

// Write content on the connection
io.WriteString(conn, fmt.Sprintf("Coucou %d\n", i))