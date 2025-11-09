package user

import (
	"database/sql"
	"iskra/shared/models"
)

type UserRepo struct {
	db *sql.DB
}

func New(db *sql.DB) (*UserRepo, error) {
	repo := UserRepo{db: db}
	err := repo.CreateTable()
	return &repo, err
}

func (r *UserRepo) CreateTable() error {
	// const op = "postgres.user.create_table"

	stmt := `
		CREATE TABLE IF NOT EXISTS public.users (
			id int NOT NULL,
			username varchar NOT NULL,
			"name" varchar NOT NULL,
			surname varchar NULL,
			age int NOT NULL,
			gender int NOT NULL,
			preferred_gender int NULL,
			career_type varchar NULL,
			personality_type varchar NULL,
			relationship_goal varchar NULL,
			important_values varchar NULL,
			city varchar NULL,
			career_place varchar NULL,
			music varchar NULL,
			films varchar NULL,
			hobbies varchar NULL,
			event_preferences varchar NULL,
			CONSTRAINT users_pk PRIMARY KEY (id),
			CONSTRAINT users_unique UNIQUE (username)
		);
	`
	// Дописать недостающее ^^^^^^

	_, err := r.db.Exec(stmt)
	return err
}

func (r *UserRepo) GetUser(ID int64) (models.UserDB, error) {
	stmt, err := r.db.Prepare(`
		SELECT id, username, name, surname, age, gender, preferred_gender, career_type, personality_type, relationship_goal, important_values, city, career_place, music, films, hobbies, event_preferences
		FROM users
		WHERE id = $1
	`)

	res := models.UserDB{}
	if err != nil {
		return res, err
	}

	err = stmt.QueryRow(ID).Scan(&res.ID, &res.Username, &res.Name, &res.Surname, &res.Age, &res.Gender, &res.PreferredGender, &res.CareerType, &res.PersonalityType, &res.RelationshipGoal, &res.ImportantValues, &res.City, &res.CareerPlace, &res.Music, &res.Films, &res.Hobbies, &res.EventPreferences)
	return res, err
}

func (r *UserRepo) GetAll() ([]models.UserDB, error) {
	stmt := `
		SELECT id, username, name, surname, age, gender, preferred_gender, career_type, personality_type, relationship_goal, important_values, city, career_place, music, films, hobbies, event_preferences
		FROM users
	`

	res := make([]models.UserDB, 0)

	rows, err := r.db.Query(stmt)
	if err != nil {
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
		return err
	}

	_, err = stmt.Exec(user.ID, user.Username, user.Name, user.Surname, user.Age, user.Gender, user.PreferredGender, user.CareerType, user.PersonalityType, user.RelationshipGoal, user.ImportantValues, user.City, user.CareerPlace, user.Music, user.Films, user.Hobbies, user.EventPreferences)
	return err
}

func (r *UserRepo) UpdateUser(user models.UserDB) error {
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
	return err
}
