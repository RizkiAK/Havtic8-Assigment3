package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type StatusData struct {
	Status struct {
		Water int `json:"water"`
		Wind  int `json:"wind"`
	} `json:"status"`
}

func main() {
	go autoReloadJson()
	http.HandleFunc("/", autoReloadWeb)
	http.ListenAndServe(":8080", nil)
}

func autoReloadJson() {
	for {
		min := 1
		max := 21
		water := rand.Intn(max-min) + min
		wind := rand.Intn(max-min) + min

		data := StatusData{}
		data.Status.Water = water
		data.Status.Wind = wind

		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Fatal("text", err.Error())
		}

		ioutil.WriteFile("data.json", jsonData, 0644)

		if err != nil {
			log.Fatal("text", err.Error())
		}
		time.Sleep(15 * time.Second)
	}
}

func autoReloadWeb(w http.ResponseWriter, r *http.Request) {
	fileData, err := ioutil.ReadFile("data.json")

	if err != nil {
		log.Fatal("text", err.Error())
	}

	var statusData StatusData

	err = json.Unmarshal(fileData, &statusData)
	if err != nil {
		log.Fatal("error occured while unMarhsalling fomr data.json file ", err.Error())
	}

	waterVal := statusData.Status.Water
	windVal := statusData.Status.Wind

	var (
		waterStatus string
		windStatus  string
	)

	switch {
	case waterVal < 5:
		waterStatus = "aman"
	default:
		waterStatus = "siaga"
	}

	switch {
	case windVal < 5:
		windStatus = "aman"
	default:
		windStatus = "siaga"
	}
	// b, err := json.Marshal(fileData)
	// if err != nil {
	// 	log.Fatal("text", err.Error())
	// }

	data := map[string]string{
		"waterStatus": waterStatus,
		"windStatus":  windStatus,
	}

	tpl, err := template.ParseFiles("index.html")

	if err != nil {
		log.Fatal("text", err.Error())
	}

	tpl.Execute(w, data)

}
