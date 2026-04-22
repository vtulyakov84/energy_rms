FROM registry.altlinux.org/c10f/postgresql:latest

ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=postgres
ENV PGDATA=/var/lib/pgsql/data

RUN chown -R postgres:postgres /var/lib/pgsql

RUN initdb -D ${PGDATA} --locale=C --encoding=UTF8

COPY config/pg_hba.conf ${PGDATA}/pg_hba.conf
COPY config/postgresql.conf ${PGDATA}/postgresql.conf

# Меняем права на конфиги
USER root
RUN chown postgres:postgres ${PGDATA}/pg_hba.conf ${PGDATA}/postgresql.conf
USER postgres

# Запуск и настройка пароля
RUN pg_ctl -D ${PGDATA} start && \
    psql -U postgres -c "ALTER USER postgres WITH PASSWORD '${POSTGRES_PASSWORD}';" && \
    pg_ctl -D ${PGDATA} stop
    
EXPOSE 5432

HEALTHCHECK --interval=30s --timeout=5s --start-period=10s --retries=3 \
  CMD pg_isready -U ${POSTGRES_USER} || exit 1

# CMD ["sh", "-c", "postgres -D ${PGDATA}"]

# Запускаем с явным указанием конфигов
CMD ["sh", "-c", "postgres -D ${PGDATA} -c config_file=${PGDATA}/postgresql.conf -c hba_file=${PGDATA}/pg_hba.conf"]