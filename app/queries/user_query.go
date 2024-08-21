package queries

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/app/models"
)


type UserQueries struct {
	db *sqlx.DB
}

func CreateUserQueries(db *sqlx.DB) *UserQueries {
	return &UserQueries{
		db: db,
	}
}


func (q *UserQueries) GetUserByEmail(email string) (*models.User, error){
	user := models.User{}
	query := `SELECT name, email, password, team, role, avatar_url FROM users where email = $1`

	if err := q.db.Get(&user, query, email); err != nil {
		fmt.Println(err)
	}

	return &user, nil
}


func (q *UserQueries) CreateUser(user *models.User) error {

	query := `INSERT INTO users (name, email, password, team, role, avatar_url) values ($1, $2, $3, $4, $5, $6)`

	if _, err := q.db.Exec(query, user.Name, user.Email, user.Password, user.Team, user.Role, user.Avatar_url); err != nil {
		return err
	}

	return nil
}