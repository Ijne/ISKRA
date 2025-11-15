Write-Host "Starting ISKRA project (Windows version)" -ForegroundColor Green

docker-compose up

# 4. Wait for databases to start
Write-Host "Waiting for databases to start (30 sec)..." -ForegroundColor Yellow
Start-Sleep -Seconds 30

# 5. Create table in PostgreSQL
Write-Host "Creating users table in PostgreSQL..." -ForegroundColor Yellow
docker-compose exec postgres psql -U postgres -d postgres -c "
CREATE TABLE IF NOT EXISTS public.users (
    id int NOT NULL,
    username varchar NULL,
    name varchar DEFAULT 'Default_user_name' NOT NULL,
    surname varchar NULL,
    age int DEFAULT 18 NULL,
    gender int DEFAULT 0 NOT NULL,
    preferred_gender int DEFAULT 2 NOT NULL,
    career_type varchar NULL,
    personality_type varchar NULL,
    relationship_goal varchar NULL,
    important_values varchar NULL,
    city varchar NULL,
    career_place varchar NULL,
    music varchar NULL,
    films varchar NULL,
    hobbies varchar NULL,
    event_preferences varchar NULL,
    photo varchar NULL,
    CONSTRAINT users_pk PRIMARY KEY (id)
);"

# 6. Fill Memgraph with test data
Write-Host "Filling Memgraph with test data..." -ForegroundColor Yellow

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
    Write-Host "Executing $script..." -ForegroundColor Cyan
    Get-Content "init-scripts\$script" | docker-compose exec -T memgraph mgconsole
    Start-Sleep -Seconds 1
}

# 7. Export data to PostgreSQL
Write-Host "Exporting data to PostgreSQL..." -ForegroundColor Yellow
Get-Content "init-scripts/12-export-to-postgres.cql" | docker-compose exec -T memgraph mgconsole | Out-File -FilePath "exported-data.sql" -Encoding UTF8

# 8. Clean exported file
Write-Host "Cleaning exported data..." -ForegroundColor Yellow
$content = Get-Content "exported-data.sql" | Where-Object { $_ -ne "" -and $_ -ne "sql_statement" }
$content = $content -replace "^\| ", ""
$content | Out-File -FilePath "exported-data.sql" -Encoding UTF8

# 9. Import to PostgreSQL
Write-Host "Importing data to PostgreSQL..." -ForegroundColor Yellow
Get-Content "exported-data.sql" | docker-compose exec -T postgres psql -U postgres -d postgres

# 10. Start main service
Write-Host "Starting main service..." -ForegroundColor Yellow
docker-compose up server -d

# 11. Check services
Write-Host "Checking running services..." -ForegroundColor Yellow
docker-compose ps

Write-Host ""
Write-Host "Done! ISKRA is running!" -ForegroundColor Green
Write-Host ""
Write-Host "Services access:" -ForegroundColor White
Write-Host "   • Application: http://localhost:8080" -ForegroundColor White
Write-Host "   • Memgraph: localhost:7687" -ForegroundColor White
Write-Host "   • PostgreSQL: localhost:5433" -ForegroundColor White
Write-Host ""
Write-Host "Created 1000 test users!" -ForegroundColor Green