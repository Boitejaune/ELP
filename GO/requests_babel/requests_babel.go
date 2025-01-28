package requests_babel

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

func Get_book(hexagon int, wall int, shelf int, volume int) string {
	// defer wg.Done()
	requestURL := "https://libraryofbabel.info/download.cgi" // url du site pour télécharger les livres
	var req *http.Response
	var err error
	if volume < 10{
		req,err = http.PostForm(requestURL,url.Values{"hex":{strconv.Itoa(hexagon)},"wall":{strconv.Itoa(wall)},"shelf":{strconv.Itoa(shelf)},"volume":{"0"+strconv.Itoa(volume)},"page":{"1"},"title":{"startofthetext"}})  // formatage de la demande POST à l'URL posée juste au dessus avec l'adresse du livre voulu
	}else{
		req,err = http.PostForm(requestURL,url.Values{"hex":{strconv.Itoa(hexagon)},"wall":{strconv.Itoa(wall)},"shelf":{strconv.Itoa(shelf)},"volume":{strconv.Itoa(volume)},"page":{"1"},"title":{"startofthetext"}})  // formatage de la demande POST à l'URL posée juste au dessus avec l'adresse du livre voulu
	}
	if err != nil{
		fmt.Printf("erreur en demandant le livre: %s", err)
		os.Exit(1)
	}
	// fmt.Printf("Received a book!\n")
	stringres,_ := io.ReadAll(req.Body) //lecture du corps du texte
	if len(stringres) < 30 {
		return("1")
	}
	final := string(stringres)
	return(final[16:len(final)-29]) //pour enlever le titre du texte au début et l'adresse du livre à la fin avant de renvoyer le contenu du livre
}