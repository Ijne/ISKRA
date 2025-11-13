package user

import (
	"database/sql"
	"iskra/shared/models"
	"iskra/shared/storage/repos"
	"log"
)

type UserRepo struct {
	db *sql.DB
}

func New(db *sql.DB) (repos.UserRepo, error) {
	repo := UserRepo{db: db}
	err := repo.CreateTable()
	return &repo, err
}

func (r *UserRepo) CreateTable() error {
	const op = "postgres.user.create_table"

	stmt := `
		CREATE TABLE IF NOT EXISTS users
			(id integer NOT NULL,
			username character varying NOT NULL,
			name character varying NOT NULL,
			surname character varying,
			age integer NOT NULL,
			gender integer NOT NULL,
			preferred_gender integer,
			career_type character varying,
			personality_type character varying,
			relationship_goal character varying,
			important_values character varying,
			city character varying,
			career_place character varying,
			music character varying,
			films character varying,
			hobbies character varying,
			event_preferences character varying);
	`

	_, err := r.db.Exec(stmt)
	if err != nil {
		log.Printf("%s: %v", op, err)
	}
	return err
}

func (r *UserRepo) GetUser(ID int64) (models.UserDB, error) {
	const op = "postgres.user.get"
	stmt, err := r.db.Prepare(`
		SELECT id, username, name, COALESCE(surname, 'nil') as surname, age, gender, COALESCE(preferred_gender, '2') as preferred_gender, COALESCE(career_type, 'nil') as career_type, COALESCE(personality_type, 'nil') as personality_type, COALESCE(relationship_goal, 'nil') as relationship_goal, COALESCE(important_values, 'nil') as important_values, COALESCE(city, 'nil') as city, COALESCE(career_place, 'nil') as career_place, COALESCE(music, 'nil') as music, COALESCE(films, 'nil') as films, COALESCE(hobbies, 'nil') as hobbies, COALESCE(event_preferences, 'nil') as event_preferences
		FROM users
		WHERE id = $1
	`)

	res := models.UserDB{}
	if err != nil {
		log.Printf("%s: %v", op, err)
		return res, err
	}

	err = stmt.QueryRow(ID).Scan(&res.ID, &res.Username, &res.Name, &res.Surname, &res.Age, &res.Gender, &res.PreferredGender, &res.CareerType, &res.PersonalityType, &res.RelationshipGoal, &res.ImportantValues, &res.City, &res.CareerPlace, &res.Music, &res.Films, &res.Hobbies, &res.EventPreferences)
	if err != nil {
		log.Printf("%s: %v", op, err)
	}
	return res, err
}

func (r *UserRepo) GetAll() ([]models.UserDB, error) {
	const op = "postgres.user.get_all"
	stmt := `
		SELECT id, username, name, COALESCE(surname, 'nil') as surname, age, gender, COALESCE(preferred_gender, '2') as preferred_gender, COALESCE(career_type, 'nil') as career_type, COALESCE(personality_type, 'nil') as personality_type, COALESCE(relationship_goal, 'nil') as relationship_goal, COALESCE(important_values, 'nil') as important_values, COALESCE(city, 'nil') as city, COALESCE(career_place, 'nil') as career_place, COALESCE(music, 'nil') as music, COALESCE(films, 'nil') as films, COALESCE(hobbies, 'nil') as hobbies, COALESCE(event_preferences, 'nil') as event_preferences
		FROM users
	`

	res := make([]models.UserDB, 0)

	rows, err := r.db.Query(stmt)
	if err != nil {
		log.Printf("%s: %v", op, err)
		return nil, err
	}

	for rows.Next() {
		var curr models.UserDB
		rows.Scan(&curr.ID, &curr.Username, &curr.Name, &curr.Surname, &curr.Age, &curr.Gender, &curr.PreferredGender, &curr.CareerType, &curr.PersonalityType, &curr.RelationshipGoal, &curr.ImportantValues, &curr.City, &curr.CareerPlace, &curr.Music, &curr.Films, &curr.Hobbies, &curr.EventPreferences)
		res = append(res, curr)
	}

	return res, nil
}

func (r *UserRepo) CreateUser(user models.UserCreate) error {
	const op = "postgres.user.create_user"
	stmt, err := r.db.Prepare(`
		INSERT INTO users (
			id,
			username,
			name,
			surname,
			age,
			gender,
			preferred_gender,
			career_type,
			personality_type,
			relationship_goal,
			important_values,
			city,
			career_place,
			music,
			films,
			hobbies,
			event_preferences
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)
	`)
	if err != nil {
		log.Printf("%s: %v", op, err)
		return err
	}

	_, err = stmt.Exec(user.ID, user.Username, user.Name, user.Surname, user.Age, user.Gender, user.PreferredGender, user.CareerType, user.PersonalityType, user.RelationshipGoal, user.ImportantValues, user.City, user.CareerPlace, user.Music, user.Films, user.Hobbies, user.EventPreferences)
	if err != nil {
		log.Printf("%s: %v", op, err)
	}
	return err
}

func (r *UserRepo) UpdateUser(user models.UserDB) error {
	const op = "postgres.user.update_user"
	stmt, err := r.db.Prepare(`
        UPDATE users 
        SET age = $1,  
            preferred_gender = $2, 
            career_type = $3, 
            personality_type = $4, 
            relationship_goal = $5, 
            important_values = $6, 
            city = $7, 
            career_place = $8, 
            music = $9, 
            films = $10, 
            hobbies = $11, 
            event_preferences = $12
        WHERE id = $13
    `)
	if err != nil {
		log.Printf("%s: %v", op, err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		user.Age,
		user.PreferredGender,
		user.CareerType,
		user.PersonalityType,
		user.RelationshipGoal,
		user.ImportantValues,
		user.City,
		user.CareerPlace,
		user.Music,
		user.Films,
		user.Hobbies,
		user.EventPreferences,
		user.ID,
	)
	if err != nil {
		log.Printf("%s: %v", op, err)
	}
	return err
}
