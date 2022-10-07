# Vagorant-Agent-Reviews

Para levantar las instancias dentro de la topología.

```sh
docker-compose up --build
```

Luego se tendría que ejecutar la ruta que migra los datos procedentes del modelo relacional sql, con la siguiente ruta:

```bash
http://localhost:3000/
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
