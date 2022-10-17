#Se establece la imágen de apache para este caso, así mismo se establece la versión
FROM httpd:2.4 AS builder

#Se establece el directorio de trabajo dentro del contenedor
WORKDIR /Frontend

#Se copian los archivos necesarios para el proyecto, en este caso todos los archivos con extensión html 
#para que queden renderizados en el directorio htdocs dada la configuración de apache
COPY /Login /usr/local/apache2/htdocs/
COPY /Templates /usr/local/apache2/htdocs/

#Se encapula el contenedor anterior en un contenedor nuevo con la imágen de nginx
FROM nginx:latest

#Se define la ruta de trabajo para su posterior renderizado en el contenedor
COPY --from=builder /usr/local/apache2/htdocs/ /usr/share/nginx/html/

#Se establece el puerto de escucha
EXPOSE 80

#Se establece el comando para ejecutar el contenedor
ENTRYPOINT ["nginx", "-g", "daemon off;"]