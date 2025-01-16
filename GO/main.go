package main

import (
	"fmt"
	"main/livenshtein"
	"main/requests_babel"
	"main/texte"
	"sync"
	"time"
	"strings"
)

var wg sync.WaitGroup // instanciation de notre structure WaitGroup

func main() {
	debut := time.Now()

	a:=requests_babel.Get_book(0,1,1,1)
	b:="Je mange des tomates"
	
	var tab[4] string
	var tab2[4] string

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