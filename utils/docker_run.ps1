# без сети:
# docker run --rm --env-file "./config/.env" -p 8080:8080 iskra:2

# с сетью:
# docker run --rm --env-file "./config/.env" --network mynet iskra:3
docker run --rm --env-file "./config/.env" -e CONFIG_PATH="./config/deploy.yaml" --network mynet iskra:3