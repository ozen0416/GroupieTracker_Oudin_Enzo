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

	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		for _, i := range Amiibo.Amiibo[0].Name {
			tmpl.Execute(w, i)
		}

	})

	fs := http.FileServer(http.Dir("style"))
	http.Handle("/style/", http.StripPrefix("/style/", fs))

	http.ListenAndServe(":80", nil)

}
