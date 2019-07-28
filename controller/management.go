package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func (c *Controller) ServeManagementApi() {
	r := mux.NewRouter()
	s := r.PathPrefix("/management").Subrouter()
	s.HandleFunc("/add/{host}", c.addHandler).Methods("POST")
	s.HandleFunc("/remove/{host}", c.removeHandler).Methods("POST")
	s.HandleFunc("", c.listHandler).Methods("GET")
	http.Handle("/", r)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
	c.wg.Done()
}

func (c *Controller) addHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	c.mr.AddMiner(vars["host"])
	w.WriteHeader(http.StatusOK)
}

func (c *Controller) removeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	c.mr.RemoveMiner(vars["host"])
	w.WriteHeader(http.StatusOK)
}

type ListResponse struct {
	Miners []string
}

func (c *Controller) listHandler(w http.ResponseWriter, r *http.Request) {
	response := ListResponse{Miners: c.mr.Miners()}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println(err)
	}
}
