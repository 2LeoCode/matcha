package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DatabaseManager struct {
	pool *pgxpool.Pool
}

func (this *DatabaseManager) Init() error {
	host, hostOk := os.LookupEnv("PGHOST")
	port, portOk := os.LookupEnv("PGPORT")
	user, userOk := os.LookupEnv("PGUSER")
	password, passwordOk := os.LookupEnv("PGPASSWORD")
	dbname, dbnameOk := os.LookupEnv("PGDATABASE")

	missingEnv := "NONE"
	switch {
	case !hostOk:
		missingEnv = "PGHOST"
	case !portOk:
		missingEnv = "PGPORT"
	case !userOk:
		missingEnv = "PGUSER"
	case !passwordOk:
		missingEnv = "PGPASSWORD"
	case !dbnameOk:
		missingEnv = "PGDATABASE"
	}
	if missingEnv != "NONE" {
		return fmt.Errorf("Environment %s not defined", missingEnv)
	}

	dbInfo := fmt.Sprintf(
		"host=%s "+
			"port=%s "+
			"user=%s "+
			"password=%s "+
			"dbname=%s ",
		host,
		port,
		user,
		password,
		dbname,
	)

	attempts := 5

	var db *pgxpool.Pool
	for {
		tmp, err := pgxpool.New(context.Background(), dbInfo)
		if err == nil {
			db = tmp
			break
		}
		if attempts == 0 {
			return err
		}
		attempts--
		log.Println("Database connection failed, retrying...")
		time.Sleep(500 * time.Millisecond)
	}

	this.pool = db

	usersErr := this.initUsersTable()

	var err error
	switch {
	case usersErr != nil:
		err = usersErr
	}
	if err != nil {
		return err
	}

	return nil
}

func (this *DatabaseManager) Close() {
	this.pool.Close()
}
