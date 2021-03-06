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

func GETAlumno(w http.ResponseWriter, r *http.Request) {
	//Estructura de la respuesta
	response := struct {
		Alumno  DB.Alumno `json:"alumno"`
		Materia string    `json:"materia"`
	}{}

	//Variable de contexto de la template
	context := map[string]interface{}{}

	//Lectura del template
	if t, err := template.ParseFiles("Alumno/templates/alumno.html"); err == nil {
		//Llenado de la template, con defer se deja
		//la sentencia para el final de la ejecucion
		defer t.Execute(w, context)
		//Lectura del parametro legajo
		vars := mux.Vars(r)
		legajo := vars["legajo"]
		//Peticion al servidor de blockchain
		if resp, err := http.Get("http://127.0.0.1:8080/api/notes/" + legajo); err == nil {
			//Lectura del body
			if body, err := ioutil.ReadAll(resp.Body); err == nil {
				//Desearializar el contendio y agregarlo al context de la template
				if err := json.Unmarshal(body, &response); err == nil {
					context["response"] = response
				} else {
					fmt.Println(err)
				}
			} else {
				fmt.Println(err)
			}
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
}

func GETNotasAlumno(w http.ResponseWriter, r *http.Request) {
	//Estructura de la respuesta
	response := struct {
		Alumno  DB.Alumno `json:"alumno"`
		Materia string    `json:"materia"`
	}{}

	//Variable de contexto de la template
	context := map[string]interface{}{}

	//Lectura del template
	if t, err := template.ParseFiles("Alumno/templates/ver_notas.html"); err == nil {
		//Llenado de la template, con defer se deja
		//la sentencia para el final de la ejecucion
		defer t.Execute(w, context)
		//Lectura del parametro legajo
		vars := mux.Vars(r)
		legajo := vars["legajo"]
		//Peticion al servidor de blockchain
		if resp, err := http.Get("http://127.0.0.1:8080/api/notes/" + legajo); err == nil {
			//Lectura del body
			if body, err := ioutil.ReadAll(resp.Body); err == nil {
				//Desearializar el contendio y agregarlo al context de la template
				if err := json.Unmarshal(body, &response); err == nil {
					context["response"] = response
				} else {
					fmt.Println(err)
				}
			} else {
				fmt.Println(err)
			}
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
}

func GETAlumnos(w http.ResponseWriter, r *http.Request) {
	//Lectura del template
	if t, err := template.ParseFiles("Alumno/templates/login.html"); err == nil {
		//Llenado de la template, con defer se deja
		//la sentencia para el final de la ejecucion
		t.Execute(w, nil)
	}
}

func main() {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/", GETAlumnos).Methods("GET")
	r.HandleFunc("/{legajo}/notas", GETNotasAlumno).Methods("GET")
	r.HandleFunc("/{legajo}", GETAlumno).Methods("GET")

	//Configuracion del servidor
	server := &http.Server{
		Addr:           ":8081",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println("Servidor del alumno funcionando en: http://127.0.0.1:8081")
	server.ListenAndServe()
}
