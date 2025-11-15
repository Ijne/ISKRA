MATCH (p:Person)
OPTIONAL MATCH (p)-[:LIVES_IN]->(c:City)
WITH p, c.name AS city_name
RETURN 
'INSERT INTO public.users (id, username, name, age, gender, preferred_gender, career_type, personality_type, relationship_goal, important_values, city, music, films, hobbies, event_preferences, photo) VALUES (' +
    toString(p.id) + ', ' +
    '\'' + p.username + '\', ' +
    '\'' + p.name + '\', ' +
    toString(p.age) + ', ' +
    toString(p.gender) + ', ' +
    toString(p.preferred_gender) + ', ' +
    '\'' + COALESCE(p.career_type, '') + '\', ' +
    '\'' + COALESCE(p.personality_type, '') + '\', ' +
    '\'' + COALESCE(p.relationship_goal, '') + '\', ' +
    '\'' + COALESCE(p.important_values, '') + '\', ' +
    '\'' + COALESCE(city_name, '') + '\', ' +
    '\'' + COALESCE(p.music, '') + '\', ' +
    '\'' + COALESCE(p.films, '') + '\', ' +
    '\'' + COALESCE(p.hobbies, '') + '\', ' +
    '\'' + COALESCE(p.event_preferences, '') + '\', ' +
    'NULL' +
');' AS sql_statement;