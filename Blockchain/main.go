package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/axelWismer/TP_ASO/DB"
	"github.com/gorilla/mux"
)

//BD de notas
var materia DB.Materia = DB.GetDB()

func GETNotasAlumno(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//OBtener el legajo
	vars := mux.Vars(r)
	leg := vars["legajo"]
	legajo, err := strconv.Atoi(leg)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//Obtener al alumno
	alumno, ok := materia.Alumnos[legajo]
	if !ok {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	//Obtener las notas del alumno para su envio
	alumno.GETNotas(materia.Notas)
	//Serializar notas
	j, err := json.Marshal(struct {
		Alumno  DB.Alumno `json:"alumno"`
		Materia string    `json:"materia"`
	}{
		alumno,
		materia.Nombre,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		//Escribir el json
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}

func GETNotasMateria(w http.ResponseWriter, r *http.Request) {
	//Serializar notas
	if j, err := json.Marshal(materia); err == nil {
		//Escribir el json
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	} else {
		fmt.Println(err)
		//Indicar un error interno
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func POSTNoteHandler(w http.ResponseWriter, r *http.Request) {
	n := struct {
		Nota       int    `json:"nota"`
		Evaluacion string `json:"evaluacion"`
		Alumno     int    `json:"alumno"`
	}{}
	nota := DB.Nota{}
	//Decodificar el json en una estructura nota

	if err := json.NewDecoder(r.Body).Decode(&n); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println(n)

	nota.Nota = n.Nota
	if eval, ok := materia.Evaluaciones[n.Evaluacion]; ok {
		nota.Evaluacion = eval
	} else {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if al, ok := materia.Alumnos[n.Alumno]; ok {
		nota.Alumno = al
	} else {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	materia.Notas = append(materia.Notas, nota)

	//Se devuelve la nota serializada
	w.Header().Set("Content-Type", "application/json")
	//Serializar el array
	j, err := json.Marshal(nota)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
		w.Write(j)
	}
}

//func DELETENoteHandler(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	id := vars["id"]
//	delete(noteStore, id)
//	w.WriteHeader(http.StatusNoContent)
//}

func main() {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/notes", GETNotasMateria).Methods("GET")
	r.HandleFunc("/api/notes/{legajo}", GETNotasAlumno).Methods("GET")

	r.HandleFunc("/api/notes", POSTNoteHandler).Methods("POST")
	//	r.HandleFunc("/api/notes/{id}", DELETENoteHandler).Methods("DELETE")

	server := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println("Blockchain funcionando en: http://127.0.0.1:8080")

	server.ListenAndServe()
}
