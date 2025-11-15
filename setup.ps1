docker-compose up

Start-Sleep -Seconds 30

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

Get-Content "init-scripts/12-export-to-postgres.cql" | docker-compose exec -T memgraph mgconsole | Out-File -FilePath "exported-data.sql" -Encoding UTF8

$content = Get-Content "exported-data.sql" | Where-Object { $_ -ne "" -and $_ -ne "sql_statement" }
$content = $content -replace "^\| ", ""
$content | Out-File -FilePath "exported-data.sql" -Encoding UTF8

Get-Content "exported-data.sql" | docker-compose exec -T postgres psql -U postgres -d postgres

docker-compose up server -d

docker-compose ps