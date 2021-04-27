package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/axelWismer/TP_ASO/DB"
	"github.com/gorilla/mux"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	response := struct {
		Alumno  DB.Alumno `json:"alumno"`
		Materia string    `json:"materia"`
	}{}

	context := map[string]interface{}{}
	t, err := template.ParseFiles("Alumno/templates/alumno.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	//Independientemente del error muestra la template
	defer t.Execute(w, context)

	//Peticion a la blockchain
	vars := mux.Vars(r)
	legajo := vars["legajo"]
	resp, err := http.Get("http://127.0.0.1:8080/api/notes/" + legajo)
	if err != nil {
		fmt.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	//Deserializa el contenido en la variable contexto
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return
	}
	context["response"] = response

}

func main() {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/{legajo}", indexHandler).Methods("GET")

	server := &http.Server{
		Addr:           ":8081",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
