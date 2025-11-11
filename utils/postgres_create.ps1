# при общении с хостом:
# docker run --name iskra-postgres -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:13.22-alpine3.22

# при общении с контейнером:
docker run --name iskra-postgres-2 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d --network mynet postgres:13.22-alpine3.22
# docker start iskra-postgres
# psql -U postgres -d postgres -p 5432