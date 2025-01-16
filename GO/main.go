package main

import (
	"fmt"
	"main/livenshtein"
	"sync"
	"time"
	"strings"
)

var wg sync.WaitGroup // instanciation de notre structure WaitGroup

func main() {
	debut := time.Now()
	a:="Je mange des tomates"
	b:="Je man, des totes mr"

	var tab[] string
	var tab2[] string

	wg := &sync.WaitGroup{}
	tab = strings.Fields(a)

	fmt.Println(a)
	fmt.Println(tab)
	var mot int = len(tab) // nombre de mots sur une page
	tab2 = strings.Fields(b)
	fmt.Println(b)
	fmt.Println(tab2)

	var res int = 0
	for i:=0 ; i<mot ; i++ {
		wg.Add(1)
		go func(){
			res += livenshtein.Livenshtein(wg,tab[i],tab2[i])}()
	}
	wg.Wait()
	fmt.Println(res)
    fin := time.Now()
    fmt.Println(fin.Sub(debut))
}