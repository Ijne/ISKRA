// Экспорт пользователей в SQL формат для PostgreSQL
MATCH (p:Person)
OPTIONAL MATCH (p)-[:LIVES_IN]->(c:City)
WITH p, c.name AS city_name
WITH 
'INSERT INTO public.users (id, username, name, age, gender, preferred_gender, career_type, personality_type, relationship_goal, important_values, city, music, films, hobbies, event_preferences, photo) VALUES (' +
    toString(p.id) + ', ' +
    '\'' + replace(p.username, '\'', '\'\'') + '\', ' +
    '\'' + replace(p.name, '\'', '\'\'') + '\', ' +
    toString(p.age) + ', ' +
    toString(p.gender) + ', ' +
    toString(p.preferred_gender) + ', ' +
    '\'' + replace(COALESCE(p.career_type, ''), '\'', '\'\'') + '\', ' +
    '\'' + replace(COALESCE(p.personality_type, ''), '\'', '\'\'') + '\', ' +
    '\'' + replace(COALESCE(p.relationship_goal, ''), '\'', '\'\'') + '\', ' +
    '\'' + replace(COALESCE(p.important_values, ''), '\'', '\'\'') + '\', ' +
    '\'' + replace(COALESCE(city_name, ''), '\'', '\'\'') + '\', ' +
    '\'' + replace(COALESCE(p.music, ''), '\'', '\'\'') + '\', ' +
    '\'' + replace(COALESCE(p.films, ''), '\'', '\'\'') + '\', ' +
    '\'' + replace(COALESCE(p.hobbies, ''), '\'', '\'\'') + '\', ' +
    '\'' + replace(COALESCE(p.event_preferences, ''), '\'', '\'\'') + '\', ' +
    'NULL' +
');' AS sql_statement
RETURN sql_statement;