# TP_ASO
## Trabajo practico integrador de la materia arquitectura de software (ASO)

Se muestra aqui un prototipo que representa una version simplificada de la arquitectura expuesta en el trabajo practico, orientada a mostrar la conexion de servicios por medio de API REST con la blockchain (representada como un microservicio en go)

## Aplicacion del alumno
### Login
![alt text](/doc/interfaces/alumno_login.png)
### Pagina de inicio
![alt text](/doc/interfaces/alumno_home.png)
### Conultar notas de la materia 
![alt text](/doc/interfaces/alumno_notas.png)

## Aplicacion del profesor
### Conultar alumnos de una materia
![alt text](/doc/interfaces/profesor_home.png)
### Conultar notas de un alumno
![alt text](/doc/interfaces/profesor_nota_ver.png)
### Registrar una nueva nota
![alt text](/doc/interfaces/profesor_nota_crear.png)

## Blockchain
La blockchain es representada por un microservicio en go y contiene el siguente esquema de datos, definidos en DB/models.go

![alt text](/doc/esquema_datos.png)

Ademas en DB/models.go se encuentra definida la BD de prueba que contiene los datos que se visualizan en las interfaces de la aplicacion del alumno y del profesor
