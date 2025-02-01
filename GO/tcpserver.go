package main

import (
	"bufio"
	"fmt"
	"log"
	"main/second"
	"net"
	"os"
	"strings"
	"sync"
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
        fmt.Println("connection accepted: ",c.RemoteAddr().String())
        go connection(c)
    }
}


// Gestion connection
func connection( c net.Conn){
    defer c.Close()

    textclient := bufio.NewReader(c) //Buffered reader : lit et met en mémoire texte en entrée

    for {
        lignclient, err := textclient.ReadString('\n') //Lit une ligne de texte du buffer (lit jusqu'à séparation '\n')
        if err != nil {
            log.Println(err)
            os.Exit(1)
        }
        request := strings.TrimSpace(lignclient) //Renvoie lignclient sans les espaces blancs de début et de fin + \n ... 

        if request == "STOP" {  //Arret de la connection par le client
            break
        }
        wg := &sync.WaitGroup{}
        wg.Add(1)
        go func(){
            request = second.Requetes(request)
            wg.Done()
        }()
        wg.Wait()
        fmt.Print("-> ", request,"\n")

        c.Write([]byte(request + "\n"))

        // Lievenstein
        //go main(request)
    }

    c.Close()
}

