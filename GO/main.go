package main

import (
	"fmt"
	"main/livenshtein"
	"main/requests_babel"
	"strings"
	"sync"
	"time"
)


func main() {
	debut := time.Now()

	b:="je mange des tomates"
	
	var tab2[] string
	wg := &sync.WaitGroup{}
	wg1 := &sync.WaitGroup{}

	// nombre de mots sur une page
	tab2 = strings.Fields(b)
	var recherche int = len(tab2)

	fmt.Println(b)
	fmt.Println(tab2)
	for w:=1 ; w<=4 ; w++{
		for s:=1 ; s<=5; s++{
			for v:=0 ; v<=32 ; v++ {
				wg1.Add(1)
				go func(s int,v int,w int){
					defer wg1.Done()
					wg.Add(1)
					a := requests_babel.Get_book(wg,3,w,s,v)
					tab := strings.Fields(a)
					wg.Wait()
					mot := len(tab)
					for i:=0 ; i<mot-recherche ; i++{
						var res int = 0
						for j:=0 ; j<recherche ; j++{
							res += livenshtein.Livenshtein(tab[i+j],tab2[j])
						}
						if res <= 11{
							fmt.Printf("%d, %s %s %s %s\n",res, tab[i], tab[i+1], tab[i+2], tab[i+3])
						}
					}
				}(s, v, w)
			}
			wg1.Wait()
		}
	}
    fin := time.Now()
    fmt.Println(fin.Sub(debut))
}