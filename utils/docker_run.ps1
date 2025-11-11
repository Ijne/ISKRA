# без сети:
# docker run --rm --env-file "./config/.env" -p 8080:8080 iskra:2

# с сетью:
docker run --rm --env-file "./config/.env" --network mynet iskra:2

# TODO: брать путь к файлу конфигурации из переменных окржунеия + сделать другой хост (название контейнера) для postgres