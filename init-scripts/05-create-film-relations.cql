MATCH (p:Person)
WITH p, p.id AS id
WITH p, id,
     ['Comedy', 'Drama', 'Action', 'Thriller', 'Horror', 'Sci-Fi', 'Fantasy', 'Romance', 'Detective', 'Adventure', 'Anime', 'Documentary', 'Art House', 'Historical', 'Crime', 'Musical'] AS film_genres
WITH p, [genre IN film_genres WHERE rand() < 0.5 | genre] AS selected_genres
WITH p, selected_genres[0..3] AS final_genres
UNWIND final_genres AS film_name
WITH p, film_name
WHERE film_name IS NOT NULL
MATCH (f:Film {name: film_name})
CREATE (p)-[:LIKES {since: date()}]->(f)
WITH p, collect(film_name) AS film_list
SET p.films = reduce(s = '', f IN film_list | s + CASE WHEN s = '' THEN '' ELSE ',' END + f);