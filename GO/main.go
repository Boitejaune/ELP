package main

import (
    "fmt"
    "io/ioutil"
)

var wg sync.WaitGroup // instanciation de notre structure WaitGroup

func main() {
	debut := time.Now()

	wg.add(1)
	go trans("baba de pla di")
	wg.add(1)
	go livenshtein()
	
	wg.Wait()
    fin := time.Now()
    fmt.Println(fin.Sub(debut))
}