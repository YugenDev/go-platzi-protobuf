FROM postgres:10.3

COPY up.sql /docker-entrypoint-initdb.d/1.sql

EXPOSE 5432

ENV POSTGRES_PASSWORD=platzi

CMD ["postgres"]