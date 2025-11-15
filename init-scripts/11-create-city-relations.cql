// Создание связей пользователей с городами
MATCH (p:Person)
WITH p, p.id AS id
WITH p, id,
     ['Москва', 'Санкт-Петербург', 'Новосибирск', 'Екатеринбург', 'Казань', 'Нижний Новгород', 'Челябинск', 'Самара', 'Омск', 'Ростов-на-Дону', 'Уфа', 'Красноярск', 'Воронеж', 'Пермь', 'Волгоград'] AS cities
WITH p, cities[id % size(cities)] AS city_name
MATCH (c:City {name: city_name})
CREATE (p)-[:LIVES_IN {since: date()}]->(c);