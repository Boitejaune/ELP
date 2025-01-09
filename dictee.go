package main

import (
    "fmt"
    "io/ioutil"
)



func read() {

    data, err := ioutil.ReadFile("dictée.txt") // lire le fichier text.txt
    if err != nil {		// si il y a une erreur 
        fmt.Println(err)
    }

    fmt.Println(string(data)) // conversion de byte en string

}

func main() {

	

}