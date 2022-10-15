<h1 align="center">
  <img src="https://static.wikia.nocookie.net/valorant/images/8/80/Valorant_Cover_Art.jpg/revision/latest/scale-to-width-down/1200?cb=20200311224515&path-prefix=es">
  <br>
</h1>

<h1 align="center">Vagorant Agent Reviews</h2>
<h3 align="center">Go and JQuery Deploy for High Disponibility Subject</h3>
<h5 align="center">Made by Brian Castro & Joaquín Fernández.</h5>

<p align="center"><img src="https://img.shields.io/github/downloads/heym1ke/Assist/total.svg?style=for-the-badge&color=f71d51" alt="Downloads"></p>



<p align="center">
  </a>
  <a href="https://playvalorant.com/es-mx/?gclid=CjwKCAjwtKmaBhBMEiwAyINuwCqcIVfypKz4Bu7T9y1K0KjUJha_BrYx27ADRc-ay-_mKUHaoGm2FRoCdRAQAvD_BwE&gclsrc=aw.ds"><img src="https://download.zone/wp-content/uploads/2020/10/Valorant-Game-Download.png" ></a>
</p>

Para levantar las instancias dentro de la topología.

```sh
docker-compose up --build
```

Para bajar las instancias del compose
```sh
docker-compose down
```

Borrar cache en contenedores
```sh
docker system prune -a
```

Borrar cache en volumenes
```sh
docker volume rm $(docker volume ls -q)
```

---
<h2 align="center">Rutas usadas:</h5>

LOGIN [POST]:

Ruta que trae los datos del usuario.
```sh
http://localhost:3000/users3
```

Body:
```JSON
{
	"email":"ana.martinez@hotmail.com",
	"password":"1234"
}
```

Response:
```json 
{
	"ID": 0,
	"CreatedAt": "0001-01-01T00:00:00Z",
	"UpdatedAt": "2022-10-10T18:03:36.953056811Z",
	"DeletedAt": null,
	"Id": 5,
	"Email": "ana.martinez@hotmail.com",
	"Password": "1234",
	"Logged_in": 1
}
```

En caso de que el usuario no se encuentre registrado, el response vendría a ser otro, similar a lo que se verá a continuación:
```json
"Usuario no encontrado"
```


Ruta con datos del back pero ahora cobre una interfaz.

```sh
http://localhost:5000/Login.html
```

Logeados [GET]:
Ruta que trae lista de usuarios con sesión ya iniciada.

```sh
http://localhost:3000/logeados
```

Response:
```json
[
	{
		"ID": 0,
		"CreatedAt": "0001-01-01T00:00:00Z",
		"UpdatedAt": "2022-10-15T21:11:43.986303Z",
		"DeletedAt": null,
		"Id": 5,
		"Email": "ana.martinez@hotmail.com",
		"Password": "1234",
		"Logged_in": 1
	}
]
```

La ruta en el cliente es aplicada en la siguiente dirección:

```sh
http://localhost:5000/astra.html
```

Mi_reseña [POST]:

Ruta usada para traer la resenia del usuario que inicio sesión respecto al agente que haya seleccionado en el menu.
```sh
http://localhost:3000/miresenia
```
Body, para este caso es necesario tanto atributo de nombre de agente y correo electrónico.
```json
{
	"name":"ASTRA",
	"email":"pedro2.rodriguez@hotmail.com"
}
```
En donde el response le responderá en caso de que se encuentre la resenia ya registrada o no.

En caso de no estar, el response vendría a ser el siguiente:
```json
"No se encontro Resenia"
```
Caso que el usuario ya haya registrado alguna reseña, el response vendría a ser el siguiente:
```json
{
	"ID": 0,
	"CreatedAt": "0001-01-01T00:00:00Z",
	"UpdatedAt": "0001-01-01T00:00:00Z",
	"DeletedAt": null,
	"Id": 1,
	"Email": "maria.gomez@gmail.com",
	"Name": "ASTRA",
	"Comment": "Mi main es la mejor del juego nasheee"
}
```
La ruta de Mi_Resenia también tiene un uso dentro de la interfaz o cliente, mediante al ruta de Astra ya mencionada (http://localhost:5000/astra.html).

All_resenias [POST]:

Ruta usada para traer todas la resenias que no son pertenecientes al usuario pero que pertenecen al mismo Agente el cuál se encuentre situado.

```sh
http://localhost:3000/allresenias
```

