package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type AmiiboList struct {
	Amiibo []struct {
		GameSeries string `json:"gameSeries"`
		Image      string `json:"image"`
		Name       string `json:"name"`
	} `json:"amiibo"`
}

func main() {

	res, err := http.Get("https://amiiboapi.com/api/amiibo/")
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var Amiibo AmiiboList

	json.Unmarshal(body, &Amiibo)

	templates := []string{
		"templates/index.html",
		"templates/header.html",
		"templates/cards.html"}

	tmpl := template.Must(template.ParseFiles(templates...))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, Amiibo.Amiibo)
	})

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.ListenAndServe(":80", nil)

}
