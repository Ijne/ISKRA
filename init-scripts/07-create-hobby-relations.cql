MATCH (p:Person)
WITH p, p.id AS id
WITH p, id,
     ['Sports', 'Travel', 'Cooking', 'Photography', 'Drawing', 'Dancing', 'Yoga', 'Cycling', 'Gaming', 'Reading', 'Gardening', 'Handicrafts', 'Music', 'Theater', 'Cinema', 'Board Games', 'Fishing', 'Hunting', 'Cars', 'Technology'] AS hobby_list
WITH p, [hobby IN hobby_list WHERE rand() < 0.5 | hobby] AS selected_hobbies
WITH p, selected_hobbies[0..3] AS final_hobbies
UNWIND final_hobbies AS hobby_name
WITH p, hobby_name
WHERE hobby_name IS NOT NULL
MATCH (h:Hobby {name: hobby_name})
CREATE (p)-[:HAS_HOBBY {since: date()}]->(h)
WITH p, collect(hobby_name) AS hobbies_list
SET p.hobbies = reduce(s = '', h IN hobbies_list | s + CASE WHEN s = '' THEN '' ELSE ',' END + h);