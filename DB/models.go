package DB

type Nota struct {
	Nota       int        `json:"nota"`
	Evaluacion Evaluacion `json:"evaluacion"`
	Alumno     Alumno     `json:"alumno"`
}

type Notas []Nota

type Alumno struct {
	Legajo   int    `json:"legajo"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Año      int    `json:"año"`
	Notas    []Nota `json:"notas"`
}

func (a *Alumno) GETNotas(notas Notas) {
	for _, n := range notas {
		if a.Legajo == n.Alumno.Legajo {
			a.Notas = append(a.Notas, n)
		}
	}
}

type Materia struct {
	Nombre       string                `json:"nombre"`
	Año          int                   `json:"año"`
	Docentes     map[int]Docente       `json:"docente"`
	Notas        Notas                 `json:"notas"`
	Alumnos      map[int]Alumno        `json:"alumnos"`
	Evaluaciones map[string]Evaluacion `json:"evaluaciones"`
}

type Docente struct {
	Legajo   int    `json:"legajo"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
}

type CategoriaEvaluacion int

const (
	Parcial CategoriaEvaluacion = iota
	TrabajoPractico
	Presentacion
)

//Permite que las categorias de evaluacion se muestren por su nombre y no por su valor
func (me CategoriaEvaluacion) String() string {
	return [...]string{"Parcial", "Trabajo Practico", "Presentación"}[me]
}

type Evaluacion struct {
	Nombre    string              `json:"nombre"`
	Categoria CategoriaEvaluacion `json:"categoria"`
	Notas     Notas               `json:"notas"`
}

func (e *Evaluacion) GETNotas(notas Notas) {
	for _, n := range notas {
		if e.Nombre == n.Evaluacion.Nombre {
			e.Notas = append(e.Notas, n)
		}
	}
}

func GetDB() Materia {

	evaluaciones := map[string]Evaluacion{
		"Entrega 1 TP":    {Nombre: "Entrega 1 TP", Categoria: CategoriaEvaluacion(TrabajoPractico)},
		"Parcial Teorico": {Nombre: "Parcial Teorico", Categoria: CategoriaEvaluacion(Parcial)},
		"Presentacion":    {Nombre: "Presentacion", Categoria: CategoriaEvaluacion(Presentacion)},
	}

	alumnos := map[int]Alumno{
		75930: {
			Legajo:   75930,
			Nombre:   "Axel",
			Apellido: "Wismer",
			Año:      2017,
		},
		75428: {
			Legajo:   75428,
			Nombre:   "Mateo",
			Apellido: "Bossio",
			Año:      2017,
		},
		75527: {
			Legajo:   75527,
			Nombre:   "Diego",
			Apellido: "Morardo",
			Año:      2017,
		},
	}

	notas := Notas{
		Nota{Nota: 8, Evaluacion: evaluaciones["Entrega 1 TP"], Alumno: alumnos[75930]},
		Nota{Nota: 9, Evaluacion: evaluaciones["Parcial Teorico"], Alumno: alumnos[75930]},
		Nota{Nota: 10, Evaluacion: evaluaciones["Presentacion"], Alumno: alumnos[75930]},
		Nota{Nota: 9, Evaluacion: evaluaciones["Entrega 1 TP"], Alumno: alumnos[75428]},
		Nota{Nota: 9, Evaluacion: evaluaciones["Parcial Teorico"], Alumno: alumnos[75428]},
		Nota{Nota: 10, Evaluacion: evaluaciones["Presentacion"], Alumno: alumnos[75428]},
		Nota{Nota: 10, Evaluacion: evaluaciones["Entrega 1 TP"], Alumno: alumnos[75527]},
		Nota{Nota: 9, Evaluacion: evaluaciones["Parcial Teorico"], Alumno: alumnos[75527]},
		Nota{Nota: 10, Evaluacion: evaluaciones["Presentacion"], Alumno: alumnos[75527]},
	}

	docentes := map[int]Docente{
		26750: {
			Legajo:   26750,
			Nombre:   "Pablo Sebastián",
			Apellido: "Frias",
		},
		17353: {
			Legajo:   17353,
			Nombre:   "Germán Ariel",
			Apellido: "Romani",
		},
	}

	materia := Materia{
		Nombre:       "Arquitectura de Software",
		Año:          2017,
		Docentes:     docentes,
		Evaluaciones: evaluaciones,
		Alumnos:      alumnos,
		Notas:        notas,
	}
	return materia
}
