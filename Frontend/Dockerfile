FROM httpd:2.4 AS builder

WORKDIR /Frontend

COPY /Login /usr/local/apache2/htdocs/

COPY /Templates /usr/local/apache2/htdocs/

FROM httpd:2.4

COPY --from=builder  /Frontend  /var/cache/httpd

COPY --from=builder  /usr/local/  /usr/local/

COPY --from=builder  /usr/local/apache2/htdocs/   /usr/local/apache2/htdocs/

EXPOSE 80

CMD ["httpd","-D" ,"FOREGROUND"]