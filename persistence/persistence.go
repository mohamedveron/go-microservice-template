package persistence

import "database/sql"

// DBSettings contains the settings of the DB
type DBSettings struct {
	Host     string
	Port     int
	Username string
	Password string
	DbName   string
}

type Persistence struct {
	dbSettings *DBSettings
	database   *sql.DB
}

func NewPersistence(dbSettings *DBSettings) *Persistence{

	return &Persistence{
		dbSettings: dbSettings,
		database: nil,
	}
}