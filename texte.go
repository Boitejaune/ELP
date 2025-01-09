package main

import (
	"fmt"
	"strings"
)

func trans(text string) [4]string {
	var tab[4] string
	var mot string
	j:=0

	for j<3 {
		for i:=0 ; i<len(text);i++ {
			if strings.Compare(string(text[i])," ")==0{
				tab[j]=mot
				mot=""
				j++
			} else{
				mot=mot+string(text[i])
				
			}
		}
	}
	tab[3]=mot
	return tab
}

func main(){
	a:="Je mange des tomates"
	tab:=trans(a)
	fmt.Println(a)
	fmt.Println(tab)
}