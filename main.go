package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func handler(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("hdwallet", "g", "--symbol", r.URL.Query().Get("symbol"))

	stdout, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, string(stdout))
}

func main() {
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
