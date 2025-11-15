#!/bin/bash

echo "🚀 Запуск ISKRA проекта (Bash версия)"

# 3. Запускаем базы данных
echo "📦 Запускаем базы данных..."
docker-compose up -d

# 4. Ждем запуска баз
echo "⏳ Ждем запуска баз данных (5 сек)..."
sleep 5

# 6. Заполняем Memgraph тестовыми данными
echo "🎯 Заполняем Memgraph тестовыми данными..."

SCRIPTS=(
    "01-create-users.cql"
    "02-create-music-nodes.cql" 
    "03-create-music-relations.cql"
    "04-create-film-nodes.cql"
    "05-create-film-relations.cql"
    "06-create-hobby-nodes.cql"
    "07-create-hobby-relations.cql"
    "08-create-event-nodes.cql"
    "09-create-event-relations.cql"
    "10-create-city-nodes.cql"
    "11-create-city-relations.cql"
)

for script in "${SCRIPTS[@]}"; do
    echo "📝 Выполняем $script..."
    cat "init-scripts/$script" | docker-compose exec -T memgraph mgconsole
    sleep 1
done

# 7. Экспортируем данные в PostgreSQL
echo "📤 Экспортируем данные в PostgreSQL..."
cat init-scripts/12-export-to-postgres.cql | docker-compose exec -T memgraph mgconsole > exported-data.sql

# 8. Чистим файл
echo "🧹 Чистим экспортированные данные..."
sed -i '/^$/d' exported-data.sql
sed -i '/^sql_statement$/d' exported-data.sql
sed -i 's/^| //g' exported-data.sql

# 9. Импортируем в PostgreSQL
echo "📥 Импортируем данные в PostgreSQL..."
docker-compose exec -T postgres psql -U postgres -d postgres -f exported-data.sql

# 10. Запускаем основной сервис
echo "🚀 Запускаем основной сервис..."
docker-compose up server -d

# 11. Финальная проверка
echo "✅ Проверяем запущенные сервисы..."
docker-compose ps

echo ""
echo "🎉 Готово! ISKRA запущена!"
echo ""
echo "📊 Доступ к сервисам:"
echo "   • Приложение: http://localhost:8080"
echo "   • Memgraph: localhost:7687" 
echo "   • PostgreSQL: localhost:5433"
echo ""
echo "👥 Создано 1000 тестовых пользователей!"