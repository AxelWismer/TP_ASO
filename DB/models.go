package DB

type Nota struct {
	Nota       int        `json:"nota"`
	Evaluacion Evaluacion `json:"evaluacion"`
}

type Notas []Nota

type Alumno struct {
	Legajo   int    `json:"legajo"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Año      int    `json:"año"`
	Notas    []Nota `json:"notas"`
}

type Materia struct {
	Nombre       string                `json:"nombre"`
	Año          int                   `json:"año"`
	Docentes     map[int]Docente       `json:"docente"`
	Notas        map[string]Nota       `json:"notas"`
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
			Notas: Notas{
				Nota{Nota: 8, Evaluacion: evaluaciones["Entrega 1 TP"]},
				Nota{Nota: 9, Evaluacion: evaluaciones["Parcial Teorico"]},
				Nota{Nota: 10, Evaluacion: evaluaciones["Presentacion"]},
			},
		},
		75428: {
			Legajo:   75428,
			Nombre:   "Mateo",
			Apellido: "Bossio",
			Año:      2017,
			Notas: Notas{
				Nota{Nota: 9, Evaluacion: evaluaciones["Entrega 1 TP"]},
				Nota{Nota: 9, Evaluacion: evaluaciones["Parcial Teorico"]},
				Nota{Nota: 10, Evaluacion: evaluaciones["Presentacion"]}},
		},
		75527: {
			Legajo:   75527,
			Nombre:   "Diego",
			Apellido: "Morardo",
			Año:      2017,
			Notas: Notas{
				Nota{Nota: 10, Evaluacion: evaluaciones["Entrega 1 TP"]},
				Nota{Nota: 9, Evaluacion: evaluaciones["Parcial Teorico"]},
				Nota{Nota: 10, Evaluacion: evaluaciones["Presentacion"]}},
		},
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
	}
	return materia
}
