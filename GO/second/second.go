package second

import (
	"fmt"
	"main/livenshtein"
	"math"
	"os"
	"strconv"
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
	var min int = 10000
	var mu sync.Mutex

	var reshex int
	var reswall int
	var resshelf int
	var resvol int
	var respage int
	var resmot string
	var breaking = false

	var hex = 1
	var w = 1
	var s = 1
	var v = 1

	for (hex<11 && !breaking) {
		//récupération du fichier texte entier de la salle
		dir := fmt.Sprintf("/mnt/d/dossiers/insa/dataset ELP/hex%d",hex)
		f,_ := os.ReadFile(dir)
		w = 1
		for (w<=4 && !breaking) {
			s = 1
			for (s<=5 && !breaking){
				v = 1
				for (v<=32 && !breaking){
					wg1.Add(1)
					go func(s int,v int,w int,hex int){
						defer wg1.Done()
						//a := requests_babel.Get_book(hex,w,s,v) //ancienne méthode avec des demandes http
						a := f[1328808*(v-1)+1328808*5*(s-1)+1328808*5*4*(w-1):1328808*v+1328808*5*(s-1)+1328808*5*4*(w-1)] //creer un livre entier composé de 410 pages de 3200 charactères
						nmb_char_passed := 0
						tab := strings.Fields(string(a)) //transformer les caractères en liste de mots
						nmb_mots := len(tab)
						for i:=0 ; i<nmb_mots-recherche ; i++{
							var res int = 0
							for j:=0 ; j<recherche ; j++{ //au cas ou le string recherché est plus grand qu'un seul mot
								res += livenshtein.Livenshtein(tab[i+j],tab2[j])
							}
							if res <= 2{
								temp := ""
								for j:=0 ; j<recherche ; j++{
									temp += tab[i+j]
								}
								// fmt.Printf("%d, %s\n", res, temp)
							}
							mu.Lock()
							if res < min{
								reshex = hex
								reswall = w
								resshelf = s
								resvol = v
								respage = int(math.Floor(410/float64(len(string(a)))*float64(nmb_char_passed)))+1 //estimation de la page de notre mot
								resmot = ""
								for j:=0 ; j<recherche ; j++{ //au cas ou le string recherché est plus grand qu'un seul mot
									resmot += tab[i+j]
								}
								min = res
							}
							mu.Unlock()
							temp := nmb_char_passed
							for string(a[temp+1]) == " " || string(a[temp+1]) == "\n"{
								temp += 1
							}
							nmb_char_passed = temp + len(tab[i])
							
						}
					}(s, v, w, hex)
					v+=1
				}
				wg1.Wait()
				if min == 0{
					breaking = true
				}
				s+=1
				// time.Sleep(5000000000) //aussi l'ancienne méthode avec les requêtes http, on devait limiter nos envois de requêtes pour éviter de faire crasher le site
			}
			w+=1
		}
		fmt.Printf("hex finished: %d\n", hex)
		hex+=1
	}
    fin := time.Now()
    fmt.Println(fin.Sub(debut))
	return("hex: " + strconv.Itoa(reshex) + " wall: " + strconv.Itoa(reswall) + " shelf: " + strconv.Itoa(resshelf) + " volume: " + strconv.Itoa(resvol) + " page environ: " + strconv.Itoa(respage) + " mot trouvé: " + resmot + " en: " + fin.Sub(debut).String())
}