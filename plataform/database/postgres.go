package database

import (
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib" // load pgx driver for PostgreSQL
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/app/queries"
)

type Queries struct {
	*queries.UserQueries
}


func OpenDBConnection() (*Queries, error){
	databaseUrl := connectionURLBuilder() 

	db, err := sqlx.Connect("pgx", databaseUrl)

	if err != nil {
		fmt.Printf("Could not connect to the database: %s", err);
	}

	return &Queries{
		UserQueries: &queries.UserQueries{db},
	}, nil

}


func connectionURLBuilder() string {
	godotenv.Load()
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_NAME"),
	)
}