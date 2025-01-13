package main

import (
    "fmt"
	"main/texte"
	"time"
	"main/livenshtein"
	"sync"
)

var wg sync.WaitGroup // instanciation de notre structure WaitGroup

func main() {
	debut := time.Now()
	a:="Je mange des tomates"
	b:="Je mange des tomates"

	wg.Add(1)
	tab:= texte.Trans(a)
	fmt.Println(a)
	fmt.Println(tab)
	var mot int = len(tab) // nombre de mots sur une page

	wg.Add(1)
	tab2:= texte.Trans(b)
	fmt.Println(b)
	fmt.Println(tab)

	var res int = 0
	for i:=0 ; i<mot ; i++ {
		wg.Add(1)
		go livenshtein.Livenshtein(tab[i],tab2[i],res)
	}
	print(res)
	wg.Wait()
    fin := time.Now()
    fmt.Println(fin.Sub(debut))
}