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
        os.Exit(1)
    }
    defer l.Close()

    for {
        c, err := l.Accept() // Accept a new connection on that port
        if err != nil {
            log.Fatal(err)
        }
        go connection(c)
    }
}


// Gestion connection
func connection( c net.Conn){
    defer c.Close()

    textclient := bufio.NewReader(c) //Buffered reader : lit et met en mémoire texte en entrée

    for {
        lignclient, err := connReader.ReadString('\n') //Lit une ligne de texte du buffer (lit jusqu'à séparation '\n')
        if err != nil {
            log.Println(err)
            os.Exit(1)
        }
        request := strings.TrimSpace(lignclient) //Renvoie lignclient sans les espaces blancs de début et de fin + \n ... 
    
        if request == "STOP" {  //Arret de la connection par le client
            break
        }
    }

    c.Close()
}


/*
// Accept a new connection on that port
conn, errconn := ln.Accept()
// Close the connection of a client
conn.Close()

// Write content on the connection
io.WriteString(conn, fmt.Sprintf("Coucou %d\n", i))


c.RemoteAddr().String() //address client
*/
