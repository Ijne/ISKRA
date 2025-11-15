MATCH (p:Person)
WITH p, p.id AS id
WITH p, id,
     ['Pop', 'Rock', 'Hip-hop', 'Electronic', 'Jazz', 'Classical', 'R&B', 'Metal', 'Indie', 'Folk', 'Country', 'Reggae', 'Blues', 'Soul', 'Disco', 'Alternative', 'Rap'] AS music_genres
WITH p, [genre IN music_genres WHERE rand() < 0.5 | genre] AS selected_genres
WITH p, selected_genres[0..3] AS final_genres
UNWIND final_genres AS music_name
WITH p, music_name
WHERE music_name IS NOT NULL
MATCH (m:Music {name: music_name})
CREATE (p)-[:LIKES {since: date()}]->(m)
WITH p, collect(music_name) AS music_list
SET p.music = reduce(s = '', m IN music_list | s + CASE WHEN s = '' THEN '' ELSE ',' END + m);