package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/PuerkitoBio/goquery"
	"encoding/json"
)

func main() {
	imdb:= getimdb("tt0068646")
	fmt.Println(imdb)

}

func getimdb(id string) string{
	url := "https://www.imdb.com/title/"+id+"/"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	
	defer resp.Body.Close()


	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	title := doc.Find("h1[data-testid=\"hero-title-block__title\"]").Text()
	
	
	score := doc.Find("div[data-testid=\"hero-rating-bar__aggregate-rating__score\"] span").First().Text()




	popularity := doc.Find("div[data-testid=\"hero-rating-bar__aggregate-rating__score\"]").Next().Next().First().Text()


    // get the image
	image, _ := doc.Find("img.ipc-image").Attr("src")


	 // get every cast
	cast:= doc.Find("a[data-testid=\"title-cast-item__actor\"]")
	for i := range cast.Nodes {
		actors = append(actors, cast.Eq(i).Text())
	}
 

	k := map[string]interface{}{
		"Title": title,
		"Score": score,
		"Popularity": popularity,
		"Image": image,
		"Director": director,
		"Cast": actors,
	}

	// convert to json
	json, err := json.Marshal(k)
	if err != nil {
		fmt.Println(err)	
	}

	return string(json)



}