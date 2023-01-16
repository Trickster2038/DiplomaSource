FROM mysql:latest

ENV MYSQL_DATABASE=levels \
    MYSQL_ROOT_PASSWORD=${MYSQL_PASS}

ADD db_dump1.sql /docker-entrypoint-initdb.d

EXPOSE 3306
