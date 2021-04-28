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

type Note struct {
	Id      int       `json:"id"`
	Note    int       `json:"note"`
	Created time.Time `json:"created"`
}
type Notes []Note

//BD de notas
var noteStore = make(map[string]Note)
var noteId = 0

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
	var note Note
	//Decodificar el json en una estructura nota
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//Incrementar el id
	noteId++
	note.Id = noteId
	//Se la marca como creada
	note.Created = time.Now()
	k := strconv.Itoa(noteId)
	noteStore[k] = note

	//Se devuelve la nota serializada
	w.Header().Set("Content-Type", "application/json")
	//Serializar el array
	j, err := json.Marshal(note)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
		w.Write(j)
	}

}

func PUTNoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var noteUpdate Note
	err := json.NewDecoder(r.Body).Decode(&noteUpdate)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if note, ok := noteStore[id]; ok {
		noteUpdate.Id, _ = strconv.Atoi(id)
		noteUpdate.Created = note.Created
		delete(noteStore, id)
		noteStore[id] = noteUpdate
	}
	w.WriteHeader(http.StatusNoContent)

}
func DELETENoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	delete(noteStore, id)
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/notes", GETNotasMateria).Methods("GET")
	r.HandleFunc("/api/notes/{legajo}", GETNotasAlumno).Methods("GET")

	r.HandleFunc("/api/notes", POSTNoteHandler).Methods("POST")
	r.HandleFunc("/api/notes/{id}", PUTNoteHandler).Methods("PUT")
	r.HandleFunc("/api/notes/{id}", DELETENoteHandler).Methods("DELETE")

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
