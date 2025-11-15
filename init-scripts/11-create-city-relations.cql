MATCH (p:Person)
WITH p, p.id AS id
WITH p, id,
     ['Moscow', 'Saint Petersburg', 'Novosibirsk', 'Yekaterinburg', 'Kazan', 'Nizhny Novgorod', 'Chelyabinsk', 'Samara', 'Omsk', 'Rostov-on-Don', 'Ufa', 'Krasnoyarsk', 'Voronezh', 'Perm', 'Volgograd'] AS cities
WITH p, cities[id % size(cities)] AS city_name
MATCH (c:City {name: city_name})
CREATE (p)-[:LIVES_IN {since: date()}]->(c);