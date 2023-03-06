package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Smash struct {
	Amiibo []struct {
		AmiiboSeries string `json:"amiiboSeries"`
		Character    string `json:"character"`
		GameSeries   string `json:"gameSeries"`
		Head         string `json:"head"`
		Image        string `json:"image"`
		Name         string `json:"name"`
		Release      struct {
			Au string `json:"au"`
			Eu string `json:"eu"`
			Jp string `json:"jp"`
			Na string `json:"na"`
		} `json:"release"`
		Tail string `json:"tail"`
		Type string `json:"type"`
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
		fmt.Println(temp.Amiibo[i].Name)
	}

	if total != nil {
		fmt.Println(total)
	}

	// tmpl := template.Must(template.ParseFiles("index.html"))

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	data := Smash{
	// 		Amiibo: {
	// 			{Character: temp.Amiibo[].Character, temp.Amiibo[10].Name},
	// 		},
	// 	}
	// 	tmpl.Execute(w, data)
	// })

	// http.ListenAndServe(":80", nil)
}
