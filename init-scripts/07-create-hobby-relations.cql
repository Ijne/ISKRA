// Создание связей пользователей с хобби
MATCH (p:Person)
WITH p, p.id AS id
WITH p, id,
     ['Спорт', 'Путешествия', 'Кулинария', 'Фотография', 'Рисование', 'Танцы', 'Йога', 'Велоспорт', 'Гейминг', 'Чтение', 'Садоводство', 'Рукоделие', 'Музыка', 'Театр', 'Кино', 'Настолки', 'Рыбалка', 'Охота', 'Авто', 'Технологии'] AS hobby_list
WITH p, [hobby IN hobby_list WHERE rand() < 0.5 | hobby] AS selected_hobbies
WITH p, selected_hobbies[0..3] AS final_hobbies
UNWIND final_hobbies AS hobby_name
WITH p, hobby_name
WHERE hobby_name IS NOT NULL
MATCH (h:Hobby {name: hobby_name})
CREATE (p)-[:HAS_HOBBY {since: date()}]->(h)
WITH p, collect(hobby_name) AS hobbies_list
SET p.hobbies = reduce(s = '', h IN hobbies_list | s + CASE WHEN s = '' THEN '' ELSE ',' END + h);