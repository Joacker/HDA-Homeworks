FROM httpd:2.4 AS builder

WORKDIR /Frontend

COPY /Login /usr/local/apache2/htdocs/

COPY /Templates /usr/local/apache2/htdocs/

FROM nginx:latest

COPY --from=builder /usr/local/apache2/htdocs/ /usr/share/nginx/html/

EXPOSE 80

ENTRYPOINT ["nginx", "-g", "daemon off;"]