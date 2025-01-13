package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

func get_book(hexagon int, wall int, shelf int, volume int) string {
	requestURL := "https://libraryofbabel.info/download.cgi" // url du site pour télécharger les livres
	req,err := http.PostForm(requestURL,url.Values{"hex":{strconv.Itoa(hexagon)},"wall":{strconv.Itoa(wall)},"shelf":{strconv.Itoa(shelf)},"volume":{"0"+strconv.Itoa(volume)},"page":{"1"},"title":{"startofthetext"}})  // formatage de la demande POST à l'URL posée juste au dessus avec l'adresse du livre voulu
	if err != nil{
		fmt.Printf("erreur en demandant le livre: %s", err)
	}
	fmt.Printf("got response!\n")
	stringres,err := io.ReadAll(req.Body) //lecture du corps du texte
	final := string(stringres)
	return(final[16:len(final)-29]) //pour enlever le titre du texte au début et l'adresse du livre à la fin avant de renvoyer le contenu du livre
}