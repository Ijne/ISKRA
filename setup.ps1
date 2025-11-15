Write-Host "🚀 Запуск ISKRA проекта (Windows версия)" -ForegroundColor Green

# 1. Клонируем репозиторий
cd ISKRA

# 2. Переключаемся на ветку dev
git checkout dev

# 3. Запускаем базы данных
Write-Host "📦 Запускаем базы данных..." -ForegroundColor Yellow
docker-compose up

# 4. Ждем запуска баз
Write-Host "⏳ Ждем запуска баз данных (30 сек)..." -ForegroundColor Yellow
Start-Sleep -Seconds 30

# 6. Заполняем Memgraph тестовыми данными
Write-Host "🎯 Заполняем Memgraph тестовыми данными..." -ForegroundColor Yellow

$scripts = @(
    "01-create-users.cql",
    "02-create-music-nodes.cql", 
    "03-create-music-relations.cql",
    "04-create-film-nodes.cql",
    "05-create-film-relations.cql",
    "06-create-hobby-nodes.cql",
    "07-create-hobby-relations.cql",
    "08-create-event-nodes.cql",
    "09-create-event-relations.cql",
    "10-create-city-nodes.cql",
    "11-create-city-relations.cql"
)

foreach ($script in $scripts) {
    Write-Host "📝 Выполняем $script..." -ForegroundColor Cyan
    Get-Content "init-scripts\$script" | docker-compose exec -T memgraph mgconsole
    Start-Sleep -Seconds 1
}

# 7. Экспортируем данные в PostgreSQL
Write-Host "📤 Экспортируем данные в PostgreSQL..." -ForegroundColor Yellow
Get-Content "init-scripts/12-export-to-postgres.cql" | docker-compose exec -T memgraph mgconsole | Out-File -FilePath "exported-data.sql" -Encoding UTF8

# 8. Чистим файл
Write-Host "🧹 Чистим экспортированные данные..." -ForegroundColor Yellow
$content = Get-Content "exported-data.sql" | Where-Object { $_ -ne "" -and $_ -ne "sql_statement" }
$content = $content -replace "^\| ", ""
$content | Out-File -FilePath "exported-data.sql" -Encoding UTF8

# 9. Импортируем в PostgreSQL
Write-Host "📥 Импортируем данные в PostgreSQL..." -ForegroundColor Yellow
Get-Content "exported-data.sql" | docker-compose exec -T postgres psql -U postgres -d postgres

# 10. Запускаем основной сервис
Write-Host "🚀 Запускаем основной сервис..." -ForegroundColor Yellow
docker-compose up server -d

# 11. Финальная проверка
Write-Host "✅ Проверяем запущенные сервисы..." -ForegroundColor Yellow
docker-compose ps

Write-Host ""
Write-Host "🎉 Готово! ISKRA запущена!" -ForegroundColor Green
Write-Host ""
Write-Host "📊 Доступ к сервисам:" -ForegroundColor White
Write-Host "   • Приложение: http://localhost:8080" -ForegroundColor White
Write-Host "   • Memgraph: localhost:7687" -ForegroundColor White
Write-Host "   • PostgreSQL: localhost:5433" -ForegroundColor White
Write-Host ""
Write-Host "👥 Создано 1000 тестовых пользователей!" -ForegroundColor Green