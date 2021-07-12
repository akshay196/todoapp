package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var DATABASE = make(map[int]string)
var LASTID int

func listTodo(w http.ResponseWriter, r *http.Request) {
	l, err := json.Marshal(DATABASE)
	if err != nil {
		http.Error(w, "Error fetching data", http.StatusOK)
		return
	}
	w.Write(l)
}

func add2Todo(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse the data", http.StatusOK)
		return
	}

	vals := r.Form
	ntasks, ok := vals["tasks"]
	if !ok {
		http.Error(w, "tasks is required", http.StatusBadRequest)
		return
	}
	for _, ntask := range ntasks {
		LASTID = LASTID + 1
		DATABASE[LASTID] = ntask
		log.Printf("%q is added to the list", ntask)
	}
}

func delTodo(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse the data", http.StatusOK)
		return
	}

	vals := r.Form
	ids, ok := vals["id"]
	if !ok {
		http.Error(w, "tasks is required", http.StatusBadRequest)
		return
	}
	for _, id := range ids {
		n, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, fmt.Sprintf("Incorrect id: %s", id), http.StatusBadRequest)
			return
		}
		delete(DATABASE, n)
	}
}

func main() {
	http.HandleFunc("/list", listTodo)
	http.HandleFunc("/add", add2Todo)
	http.HandleFunc("/done", delTodo)

	log.Println("Listening on :8000")
	err := http.ListenAndServe(":8000", nil)
	log.Fatal(err)
}
