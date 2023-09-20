package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/fmhr/fj"
)

func main() {
	log.Println("starting server...")
	http.HandleFunc("/", handler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}
	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	cnf := setConf()
	// seed
	seedString := r.URL.Query().Get("seed")
	if seedString == "" {
		fmt.Fprint(w, "seed is empty")
		return
	}
	seed, err := strconv.Atoi(seedString)
	if err != nil {
		fmt.Fprintf(w, "seed is not a number: %s", err)
		return
	}
	fmt.Fprintf(w, "seed is %d", seed)
	fj.Gen(&cnf, seed) // generate in/{seed}.txt
	// Config
	// reactive
	reactiveString := r.URL.Query().Get("reactive")
	var rtn map[string]float64
	if reactiveString == "" {
		cnf.Reactive = false
		rtn, err = fj.RunVis(&cnf, seed)
	} else {
		cnf.Reactive = true
		rtn, err = fj.ReactiveRun(&cnf, seed)
	}
	if err != nil {
		http.Error(w, fmt.Sprintf("run error: %s", err), http.StatusInternalServerError)
	}
	jdonData, err := json.Marshal(rtn)
	if err != nil {
		http.Error(w, fmt.Sprintf("json error: %s", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jdonData)
}

func setConf() (c fj.Config) {
	c.GenPath = "gen"
	return
}
