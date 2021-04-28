package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/axelWismer/TP_ASO/DB"
	"github.com/gorilla/mux"
)

func GETNotasMateria(w http.ResponseWriter, r *http.Request) {
	//Variable de contexto de la template
	var materia DB.Materia
	context := map[string]interface{}{}

	//Lectura del template
	if t, err := template.ParseFiles("Profesor/templates/profesor.html"); err == nil {
		//Llenado de la template, con defer se deja
		//la sentencia para el final de la ejecucion
		defer t.Execute(w, context)
		//Peticion al servidor de blockchain
		if resp, err := http.Get("http://127.0.0.1:8080/api/notes"); err == nil {
			//Lectura del body
			if body, err := ioutil.ReadAll(resp.Body); err == nil {
				//Desearializar el contendio y agregarlo al context de la template
				if err := json.Unmarshal(body, &materia); err == nil {
					//Obtener las notas por evaluacion
					//Guardarlas nuevamente en el diccionario
					for k := range materia.Evaluaciones {
						ev := materia.Evaluaciones[k]
						ev.GETNotas(materia.Notas)
						materia.Evaluaciones[k] = ev
					}
					//Agregar la materia al contexto
					context["Materia"] = materia
				} else {
					//Se muestran todos los errores del proceso por pantalla
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

func GETCrearNota(w http.ResponseWriter, r *http.Request) {
	//Variable de contexto de la template
	var materia DB.Materia
	context := map[string]interface{}{}

	//Lectura del template
	if t, err := template.ParseFiles("Profesor/templates/crear_nota.html"); err == nil {
		//Llenado de la template, con defer se deja
		//la sentencia para el final de la ejecucion
		defer t.Execute(w, context)
		//Peticion al servidor de blockchain
		if resp, err := http.Get("http://127.0.0.1:8080/api/notes"); err == nil {
			//Lectura del body
			if body, err := ioutil.ReadAll(resp.Body); err == nil {
				//Desearializar el contendio y agregarlo al context de la template
				if err := json.Unmarshal(body, &materia); err == nil {
					//Obtener las notas por evaluacion
					//Guardarlas nuevamente en el diccionario
					//Agregar la materia al contexto
					context["Materia"] = materia
				} else {
					//Se muestran todos los errores del proceso por pantalla
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

func POSTCrearNota(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	n, _ := strconv.Atoi(r.Form["nota"][0])
	a, _ := strconv.Atoi(r.Form["alumno"][0])

	nota := struct {
		Nota       int    `json:"nota"`
		Evaluacion string `json:"evaluacion"`
		Alumno     int    `json:"alumno"`
	}{
		n,
		r.Form["evaluacion"][0],
		a,
	}
	fmt.Println(nota)
	if j, err := json.Marshal(nota); err == nil {
		//Escribir el json
		responseBody := bytes.NewBuffer(j)
		if _, err := http.Post("http://127.0.0.1:8080/api/notes", "application/json", responseBody); err == nil {
			http.Redirect(w, r, "http://127.0.0.1:8082/", http.StatusSeeOther)
			return
		}
		GETCrearNota(w, r)
	}
}

func main() {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/", GETNotasMateria).Methods("GET")
	r.HandleFunc("/nota/crear", GETCrearNota).Methods("GET")
	r.HandleFunc("/nota/crear", POSTCrearNota).Methods("POST")

	//Configuracion del servidor
	server := &http.Server{
		Addr:           ":8082",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println("Servidor del profesor funcionando en: http://127.0.0.1:8082")
	server.ListenAndServe()
}
