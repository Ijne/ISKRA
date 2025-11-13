package socialweb

import (
	"context"
	"fmt"
	"iskra/shared/models"
	"iskra/shared/storage/repos"
	"strings"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type SocialWebRepo struct {
	driver neo4j.DriverWithContext
}

func New(driver neo4j.DriverWithContext) (repos.SocialWebRepo, error) {
	repo := SocialWebRepo{driver: driver}
	return &repo, nil
}

func (r *SocialWebRepo) CreateUser(user models.UserCreate) error {
	session := r.driver.NewSession(context.Background(), neo4j.SessionConfig{})
	defer session.Close(context.Background())

	if _, err := session.ExecuteWrite(context.Background(), func(tx neo4j.ManagedTransaction) (interface{}, error) {
		query1 := `
			CREATE (p:Person {
				id: $id, 
				name: $name, 
				username: $username, 
				age: $age, 
				gender: $gender, 
				preferred_gender: $p_g, 
				career_type: $c_t, 
				personality_type: $p_t, 
				relationship_goal: $r_g, 
				important_values: $i_v,
				currentLimit: 30,
				lastLimitUpdate: toString(datetime()),
				maxLimit: 30,
				refreshAmount: 5,
				refreshRate: "PT1H",
				music: $music_string,
				films: $films_string,
				hobbies: $hobbies_string,
				event_preferences: $events_string
			})
		`
		query2 := `	
			MATCH (p:Person {id: $id})
			UNWIND $music AS music
			WITH p, music WHERE music <> 'nil'
			MERGE (m:Music {name: music})
			CREATE (p)-[:LIKES {since: date()}]->(m)
		`
		query3 := `	
			MATCH (p:Person {id: $id})
			UNWIND $hobbies AS hobby
			WITH p, hobby WHERE hobby <> 'nil'
			MERGE (h:Hobby {name: hobby})
			CREATE (p)-[:HAS_HOBBY {since: date()}]->(h)
		`
		query4 := `	
			MATCH (p:Person {id: $id})
			UNWIND $films AS film
			WITH p, film WHERE film <> 'nil'
			MERGE (f:Film {name: film})
			CREATE (p)-[:LIKES {since: date()}]->(f)
		`
		query5 := `	
			MATCH (p:Person {id: $id})
			UNWIND $events AS event
			WITH p, event WHERE event <> 'nil'
			MERGE (e:Event {name: event})
			CREATE (p)-[:WANTS_VISIT {since: date()}]->(e)
		`
		query6 := `	
			MATCH (p:Person {id: $id})
			WITH p, $city AS city WHERE city <> 'nil'
			MERGE (c:City {name: $city})
			MERGE (p)-[r:LIVES_IN]->(c)
			SET r.since = date()
		`

		params := map[string]interface{}{
			"id":             user.ID,
			"name":           user.Name,
			"username":       user.Username,
			"age":            user.Age,
			"city":           user.City,
			"gender":         user.Gender,
			"p_g":            user.PreferredGender,
			"c_t":            user.CareerType,
			"p_t":            user.PersonalityType,
			"r_g":            user.RelationshipGoal,
			"i_v":            user.ImportantValues,
			"music":          strings.Split(user.Music, ","),
			"films":          strings.Split(user.Films, ","),
			"hobbies":        strings.Split(user.Hobbies, ","),
			"events":         strings.Split(user.EventPreferences, ","),
			"music_string":   user.Music,
			"films_string":   user.Films,
			"hobbies_string": user.Hobbies,
			"events_string":  user.EventPreferences,
		}

		result, err := tx.Run(context.Background(), query1, params)
		if err != nil {
			return nil, err
		}
		_, err = tx.Run(context.Background(), query2, params)
		if err != nil {
			return nil, err
		}
		_, err = tx.Run(context.Background(), query3, params)
		if err != nil {
			return nil, err
		}
		_, err = tx.Run(context.Background(), query4, params)
		if err != nil {
			return nil, err
		}
		_, err = tx.Run(context.Background(), query5, params)
		if err != nil {
			return nil, err
		}
		_, err = tx.Run(context.Background(), query6, params)
		if err != nil {
			return nil, err
		}

		return result.Consume(context.Background())
	}); err != nil {
		return err
	}

	return nil
}

func (r *SocialWebRepo) UpdateUser(user models.UserDB) error {
	session := r.driver.NewSession(context.Background(), neo4j.SessionConfig{})
	defer session.Close(context.Background())

	if _, err := session.ExecuteWrite(context.Background(), func(tx neo4j.ManagedTransaction) (interface{}, error) {
		query1 := `
			MATCH (p:Person {id: $id})
			SET
				p.id = $id, 
				p.name = $name, 
				p.username = $username, 
				p.age = $age, 
				p.gender = $gender, 
				p.preferred_gender = $p_g, 
				p.career_type = $c_t, 
				p.personality_type = $p_t, 
				p.relationship_goal = $r_g, 
				p.important_values = $i_v
			RETURN p
		`

		query2 := `
			MATCH (p:Person {id: $id})
			OPTIONAL MATCH (p)-[r:LIKES]->(music:Music)
			WHERE NOT music.name IN $newMusic
			DELETE r
			
			WITH p, $newMusic AS targetMusic
			UNWIND targetMusic AS musicName
			OPTIONAL MATCH (p)-[:LIKES]->(existingMusic:Music {name: musicName})
			WITH p, musicName
			WHERE existingMusic IS NULL AND musicName <> "nil" AND musicName <> ""
			MERGE (music:Music {name: musicName})
			CREATE (p)-[:LIKES {since: date()}]->(music)
		`

		query3 := `
			MATCH (p:Person {id: $id})
			OPTIONAL MATCH (p)-[r:LIKES]->(film:Film)
			WHERE NOT film.name IN $newFilms
			DELETE r
			
			WITH p, $newFilms as targetFilms
			UNWIND targetFilms AS filmName
			OPTIONAL MATCH (p)-[:LIKES]->(existingFilm:Film {name: filmName})
			WITH p, filmName
			WHERE existingFilm IS NULL AND filmName <> "nil" AND filmName <> ""
			MERGE (film:Film {name: filmName})
			CREATE (p)-[:LIKES {since: date()}]->(film)
		`

		query4 := `
			MATCH (p:Person {id: $id})
			OPTIONAL MATCH (p)-[r:HAS_HOBBY]->(hobby:Hobby)
			WHERE NOT hobby.name IN $newHobbies
			DELETE r
			
			WITH p, $newHobbies as targetHobbies
			UNWIND targetHobbies AS hobbyName
			OPTIONAL MATCH (p)-[:HAS_HOBBY]->(existingHobby:Hobby {name: hobbyName})
			WITH p, hobbyName
			WHERE existingHobby  IS NULL AND hobbyName <> "nil" AND hobbyName <> "" 
			MERGE (hobby:Hobby {name: hobbyName})
			CREATE (p)-[:HAS_HOBBY {since: date()}]->(hobby)
		`

		query5 := `
			MATCH (p:Person {id: $id})
			OPTIONAL MATCH (p)-[r:WANTS_VISIT]->(event:Event)
			WHERE NOT event.name IN $newEvents
			DELETE r
			
			WITH p, $newEvents as targetEvents
			UNWIND targetEvents AS eventName
			OPTIONAL MATCH (p)-[:WANTS_VISIT]->(existingEvent:Event {name: eventName})
			WITH p, eventName
			WHERE existingEvent IS NULL AND eventName <> "nil" AND eventName <> ""
			MERGE (event:Event {name: eventName})
			CREATE (p)-[:WANTS_VISIT {since: date()}]->(event)
		`

		query6 := `
			MATCH (p:Person {id: $id})
			OPTIONAL MATCH (p)-[r:LIVES_IN]->(:City)
			DELETE r
			WITH p
			WHERE $city <> "nil" AND $city <> ""
			MERGE (newCity:City {name: $city})
			CREATE (p)-[:LIVES_IN {since: date()}]->(newCity)
		`

		params := map[string]interface{}{
			"id":         user.ID,
			"name":       user.Name,
			"username":   user.Username,
			"age":        user.Age,
			"city":       user.City,
			"gender":     user.Gender,
			"p_g":        user.PreferredGender,
			"c_t":        user.CareerType,
			"p_t":        user.PersonalityType,
			"r_g":        user.RelationshipGoal,
			"i_v":        user.ImportantValues,
			"newMusic":   strings.Split(user.Music, ","),
			"newFilms":   strings.Split(user.Films, ","),
			"newHobbies": strings.Split(user.Hobbies, ","),
			"newEvents":  strings.Split(user.EventPreferences, ","),
		}

		result, err := tx.Run(context.Background(), query1, params)
		if err != nil {
			return nil, err
		}

		_, err = tx.Run(context.Background(), query2, params)
		if err != nil {
			return nil, err
		}
		_, err = tx.Run(context.Background(), query3, params)
		if err != nil {
			return nil, err
		}
		_, err = tx.Run(context.Background(), query4, params)
		if err != nil {
			return nil, err
		}
		_, err = tx.Run(context.Background(), query5, params)
		if err != nil {
			return nil, err
		}
		_, err = tx.Run(context.Background(), query6, params)
		if err != nil {
			return nil, err
		}

		return result.Consume(context.Background())
	}); err != nil {
		return err
	}
	return nil
}

func (r *SocialWebRepo) GetRecommendations(id int64) ([]models.UserResponse, error) {
	session := r.driver.NewSession(context.Background(), neo4j.SessionConfig{})
	defer session.Close(context.Background())

	recommendations, err := session.ExecuteRead(context.Background(), func(tx neo4j.ManagedTransaction) (interface{}, error) {
		params := map[string]interface{}{
			"id":              id,
			"tier1_min_score": 5,
		}

		limit := int64(10)
		limit_query := `MATCH (me:Person {id: $id}) RETURN me.currentLimit AS limit`
		limit_result, err := tx.Run(context.Background(), limit_query, params)
		if err == nil && limit_result.Next(context.Background()) {
			if limitValue, ok := limit_result.Record().Get("limit"); ok {
				if intValue, ok := limitValue.(int64); ok {
					limit = intValue
				}
			}
		}
		params["limit"] = limit

		query := `
            MATCH (me:Person {id: $id})

            MATCH (candidate:Person)
            WHERE candidate.id <> me.id
            AND (candidate.gender = me.preferred_gender OR me.preferred_gender = 2)
            AND (me.gender = candidate.preferred_gender OR candidate.preferred_gender = 2)

            OPTIONAL MATCH (me)-[swipe:SWIPED]->(candidate)
            WITH candidate, me, swipe
            WHERE swipe IS NULL

            OPTIONAL MATCH (me)-[:LIVES_IN]->(city:City)<-[:LIVES_IN]-(candidate)
            OPTIONAL MATCH (me)-[:WANTS_VISIT]->(event:Event)<-[:WANTS_VISIT]-(candidate)
            OPTIONAL MATCH (me)-[:LIKES]->(film:Film)<-[:LIKES]-(candidate)
            OPTIONAL MATCH (me)-[:LIKES]->(music:Music)<-[:LIKES]-(candidate)

            WITH candidate, me, city,
                (COUNT(DISTINCT event) * 10) AS event_score,
                (CASE WHEN city IS NOT NULL THEN 15 ELSE 0 END) AS city_score,
                (COUNT(DISTINCT film) + COUNT(DISTINCT music)) * 2 AS taste_score,
                -(abs(me.age - candidate.age) * 0.5) AS age_penalty

            WITH candidate, city, (event_score + city_score + taste_score + age_penalty) AS total_score
            WHERE total_score >= -100000

            RETURN 
                candidate.id AS id,
                candidate.username AS username,
                candidate.name AS name,
                candidate.surname AS surname,
                candidate.age AS age,
                candidate.gender AS gender,
                candidate.preferred_gender AS preferred_gender,
                candidate.career_type AS career_type,
                candidate.personality_type AS personality_type,
                candidate.relationship_goal AS relationship_goal,
                candidate.important_values AS important_values,
                city.name AS city,
                candidate.career_place AS career_place,
                candidate.music AS music,
                candidate.films AS films,
                candidate.hobbies AS hobbies,
                candidate.event_preferences AS event_preferences,
                total_score AS match_score
            ORDER BY total_score DESC
            LIMIT 50
        `

		result, err := tx.Run(context.Background(), query, params)
		if err != nil {
			return nil, err
		}

		var users []models.UserResponse
		for result.Next(context.Background()) {
			record := result.Record()
			if record == nil {
				continue
			}

			user := models.UserResponse{}

			if idVal, ok := record.Get("id"); ok {
				if idInt, ok := idVal.(int64); ok {
					user.ID = idInt
				}
				print(user.ID)
			}

			if username, ok := record.Get("username"); ok {
				user.Username, _ = username.(string)
			}

			if name, ok := record.Get("name"); ok {
				user.Name, _ = name.(string)
			}

			if age, ok := record.Get("age"); ok {
				if ageInt, ok := age.(int64); ok {
					user.Age = int(ageInt)
				}
			}

			if gender, ok := record.Get("gender"); ok {
				if genderInt, ok := gender.(int64); ok {
					user.Gender = int(genderInt)
				}
			}

			if surname, ok := record.Get("surname"); ok && surname != nil {
				if surnameStr, ok := surname.(string); ok {
					user.Surname = &surnameStr
				}
			}

			if preferredGender, ok := record.Get("preferred_gender"); ok {
				if pgInt, ok := preferredGender.(int64); ok {
					user.PreferredGender = int(pgInt)
				}
			}

			if careerType, ok := record.Get("career_type"); ok {
				user.CareerType, _ = careerType.(string)
			}

			if personalityType, ok := record.Get("personality_type"); ok {
				user.PersonalityType, _ = personalityType.(string)
			}

			if relationshipGoal, ok := record.Get("relationship_goal"); ok {
				user.RelationshipGoal, _ = relationshipGoal.(string)
			}

			if importantValues, ok := record.Get("important_values"); ok {
				user.ImportantValues, _ = importantValues.(string)
			}

			if city, ok := record.Get("city"); ok {
				user.City, _ = city.(string)
			}

			if careerPlace, ok := record.Get("career_place"); ok {
				user.CareerPlace, _ = careerPlace.(string)
			}

			if music, ok := record.Get("music"); ok {
				user.Music, _ = music.(string)
			}

			if films, ok := record.Get("films"); ok {
				user.Films, _ = films.(string)
			}

			if hobbies, ok := record.Get("hobbies"); ok {
				user.Hobbies, _ = hobbies.(string)
			}

			if eventPreferences, ok := record.Get("event_preferences"); ok {
				user.EventPreferences, _ = eventPreferences.(string)
			}

			if matchScore, ok := record.Get("match_score"); ok {
				fmt.Println(matchScore)
			}

			users = append(users, user)
		}

		if err := result.Err(); err != nil {
			return nil, err
		}

		return users, nil
	})

	if err != nil {
		return nil, err
	}

	if recs, ok := recommendations.([]models.UserResponse); ok {
		return recs, nil
	}

	return []models.UserResponse{}, nil
}

func (r *SocialWebRepo) SetSwipe(id1, id2 int64, interaction_type string) error {
	session := r.driver.NewSession(context.Background(), neo4j.SessionConfig{})
	defer session.Close(context.Background())

	if _, err := session.ExecuteWrite(context.Background(), func(tx neo4j.ManagedTransaction) (interface{}, error) {
		params := map[string]interface{}{
			"id1":              id1,
			"id2":              id2,
			"interaction_type": interaction_type,
		}

		query := `
			MATCH (u1:Person {id: $id1})
			MATCH (u2:Person {id: $id2})
			MERGE (u1)-[swipe:SWIPED]->(u2)
			SET swipe.interaction_type = $interaction_type
        `

		result, err := tx.Run(context.Background(), query, params)
		if err != nil {
			return nil, err
		}

		return result.Consume(context.Background())
	}); err != nil {
		return err
	}

	return nil
}
