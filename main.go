package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"text/template"
	"time"
)

type Smash struct {
	Amiibo []struct {
		GameSeries string `json:"gameSeries"`
		Image      string `json:"image"`
		Name       string `json:"name"`
	} `json:"amiibo"`
}

func main() {
	url := "https://amiiboapi.com/api/amiibo/"

	timeClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("User-Agent", "spacecount-total")

	res, getErr := timeClient.Do(req)
	if getErr != nil {
		fmt.Println(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		fmt.Println(readErr)
	}

	var temp Smash
	total := json.Unmarshal(body, &temp)
	for i := 0; i < len(temp.Amiibo); i++ {
		fmt.Println(temp.Amiibo[i])
	}

	if total != nil {
		fmt.Println(total)
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html", "templates/test.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, temp.Amiibo)
	})

	http.ListenAndServe(":80", nil)
}
