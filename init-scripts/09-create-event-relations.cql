MATCH (p:Person)
WITH p, p.id AS id
WITH p, id,
     ['Concerts', 'Cinema', 'Exhibitions', 'Theaters', 'Festivals', 'Sports Events', 'Parties', 'Clubs', 'Restaurants', 'Cafes', 'Picnics', 'Hiking', 'Master Classes', 'Lectures', 'Yoga Sessions', 'Dancing', 'Board Games', 'Karaoke', 'Bowling', 'Karting'] AS event_list
WITH p, [event IN event_list WHERE rand() < 0.5 | event] AS selected_events
WITH p, selected_events[0..3] AS final_events
UNWIND final_events AS event_name
WITH p, event_name
WHERE event_name IS NOT NULL
MATCH (e:Event {name: event_name})
CREATE (p)-[:WANTS_VISIT {since: date()}]->(e)
WITH p, collect(event_name) AS events_list
SET p.event_preferences = reduce(s = '', e IN events_list | s + CASE WHEN s = '' THEN '' ELSE ',' END + e);