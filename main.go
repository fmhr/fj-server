package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/fmhr/fj"
	"github.com/pelletier/go-toml/v2"
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
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var data map[string]interface{}
	if err := toml.Unmarshal(body, &data); err != nil {
		http.Error(w, "Error unmarshaling request body", http.StatusInternalServerError)
		return
	}
	log.Println("received request")

	cmd, ok := data["Cmd"].(string)
	if !ok {
		http.Error(w, "Cmd is missing or not a string", http.StatusBadRequest)
		return
	}
	fmt.Printf("Received command: %s\n", cmd)
	runCloud(0)

	w.Write([]byte("Hello " + cmd + "!\n"))
}

// tester はある。
func runCloud(seed int) {
	var cnf fj.Config
	cnf.GenPath = "./gen"
	fj.Gen(&cnf, seed)
}
