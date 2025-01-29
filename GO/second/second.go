package second

import (
	"fmt"
	"main/livenshtein"
	"main/requests_babel"
	"math/rand/v2"
	"strings"
	"sync"
	"time"
)


func Requetes(requete string) string {
	debut := time.Now()
	
	b := requete
	var tab2[] string
	wg1 := &sync.WaitGroup{}

	// nombre de mots sur une page
	tab2 = strings.Fields(b)
	var recherche int = len(tab2)
	var hex int = rand.IntN(100) // choix d'une salle au hasard dans la librairie de babel
	var min int = 10000

	var reshex int
	var reswall int
	var resshelf int
	var resvol int
	//pour des raisons 
	fmt.Println(tab2)
	for w:=1 ; w<=4 ; w++{
		for s:=1 ; s<=5; s++{
			for v:=0 ; v<=32 ; v++ {
				wg1.Add(1)
				go func(s int,v int,w int){
					defer wg1.Done()
					a := requests_babel.Get_book(hex,w,s,v)
					if a == "1" {
						fmt.Println("erreur en téléchargeant le livre")
					}else
					{
						tab := strings.Fields(a)
						mot := len(tab)
						for i:=0 ; i<mot-recherche ; i++{
							var res int = 0
							for j:=0 ; j<recherche ; j++{
								res += livenshtein.Livenshtein(tab[i+j],tab2[j])
							}
							if res <= 3{
								fmt.Printf("%d, %s\n",res, tab[i])
							}
							if res < min{
								reshex = hex
								reswall = w
								resshelf = s
								resvol = v
								min = res
							}
						}
					}
				}(s, v, w)
			}
			if min == 0{
				return("hex: " + string(reshex) + " wall: " + string(reswall) + " shelf: " + string(resshelf) + " volume: " + string(resvol))
			}
			time.Sleep(5000000000)
		}
		wg1.Wait()
	}
    fin := time.Now()
    fmt.Println(fin.Sub(debut))
	return("hex: " + string(reshex) + " wall: " + string(reswall) + " shelf: " + string(resshelf) + " volume: " + string(resvol))
}