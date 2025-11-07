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
			CONSTRAINT users_pk PRIMARY KEY (id),
			CONSTRAINT users_unique UNIQUE (username)
		);
	`

	_, err := r.db.Exec(stmt)
	return err
}

func (r *UserRepo) GetUser(ID int64) (models.UserDB, error) {
	stmt, err := r.db.Prepare(`
		SELECT id, username, name, surname, age, gender, preferred_gender, career_type, personality_type, relationship_goal, important_values, city, career_place
		FROM users
		WHERE id = $1
	`)

	res := models.UserDB{}
	if err != nil {
		return res, err
	}

	err = stmt.QueryRow(ID).Scan(&res.ID, &res.Username, &res.Name, &res.Surname, &res.Age, &res.Gender, &res.PreferredGender, &res.CareerType, &res.PersonalityType, &res.RelationshipGoal, &res.ImportantValues, &res.City, &res.CareerPlace)
	return res, err
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
			career_place
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(user.ID, user.Username, user.Name, user.Surname, user.Age, user.Gender, user.PreferredGender, user.CareerType, user.PersonalityType, user.RelationshipGoal, user.ImportantValues, user.City, user.CareerPlace)
	return err
}
