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
	
	wg1 := &sync.WaitGroup{}

	b := requete
	var tab2[] string = strings.Fields(b) //mots recherchés
	var recherche int = len(tab2) // nmb mots recherchés
	var min int = 10000
	var mu sync.Mutex

	var res[] int = make([]int, 5)
	var resmot string
	var breaking = false

	var hex = 1
	var w = 1
	var s = 1
	var v = 1

	for (hex<31 && !breaking) {
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

					//recherche du/des mot(s) à l'intérieur d'un livre, on recherche dans les 32 livres d'une étagère en même temps
					go func(s int,v int,w int,hex int){
						defer wg1.Done()
						
						//a := requests_babel.Get_book(hex,w,s,v) //ancienne méthode avec des demandes http

						//un simple check de la longueur de la base de donnée, au cas où elle est pas exacte
						var a string
						if len(f) > 1328808*v+1328808*32*(s-1)+1328808*32*5*(w-1){
							a = string(f[1328808*(v-1)+1328808*32*(s-1)+1328808*32*5*(w-1):1328808*v+1328808*32*(s-1)+1328808*32*5*(w-1)]) //creer un livre entier composé de 410 pages de 3200 charactères
						}else{
							a = ""
						}

						nmb_char_passed := 0 //on stocke le nombre de caractères qu'on a passé pour trouver le numéro de la page
						tab := strings.Fields(a) //transformer les caractères en liste de mots
						nmb_mots := len(tab)

						//boucle principale
						for i:=0 ; i<nmb_mots-recherche ; i++{
							
							var resultat int = 0
							for j:=0 ; j<recherche ; j++{ //au cas ou le string recherché est plus grand qu'un seul mot
								resultat += livenshtein.Livenshtein(tab[i+j],tab2[j]) //calcul de la distance de livenshtein avec le mot recherché et le mot actuel dans le livre
							}

							//on lock le mutex pour éviter que d'autres process modifient en même temps les résultats
							mu.Lock()
							if resultat < min{
								res[0] = hex
								res[1] = w
								res[2] = s
								res[3] = v
								res[4] = int(math.Floor(410/float64(len(string(a)))*float64(nmb_char_passed)))+1 //estimation de la page de notre mot
								resmot = ""
								for j:=0 ; j<recherche ; j++{ //au cas ou le string recherché est plus grand qu'un seul mot
									resmot += tab[i+j]
									resmot += " "
								}
								min = resultat
							}
							mu.Unlock()
							
							//calcul du nombre de charactères, qui ne sont plus dans la liste tab, déjà vérifiés
							temp := nmb_char_passed
							if a != ""{
								for string(a[temp+1]) == " " || string(a[temp+1]) == "\n"{
									temp += 1
								}
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
	return("hex: " + strconv.Itoa(res[0]) + " wall: " + strconv.Itoa(res[1]) + " shelf: " + strconv.Itoa(res[2]) + " volume: " + strconv.Itoa(res[3]) + " page environ: " + strconv.Itoa(res[4]) + " mot trouvé: " + resmot + " en: " + fin.Sub(debut).String())
}