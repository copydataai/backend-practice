FROM postgres:14.1-alpine

# For apply replication
# RUN rm -rf /var/lib/psql/data/*

# RUN pg_basebackup -U postgres -R -D /var/lib/pgsql/data/ --host=172.168.1.24 --port=5432

# COPY ./compose/postgres/Read/postgresql_example.conf /var/lib/postgresql/data/postgresql.conf
COPY ./compose/postgres/Read/setup-master.sh /docker-entrypoint-initdb.d/docker-entrypoint.sh

RUN chmod 0666 /docker-entrypoint-initdb.d/docker-entrypoint.sh
