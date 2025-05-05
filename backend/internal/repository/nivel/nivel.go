package nivel

import "database/sql"

type NivelRepository struct {
	db *sql.DB
}

const (
	// Error Messages
	errorFetchingData       string = "error fetching data"
	errorPreparingStatement string = "error preparing statement"
	errorExecutingStatement string = "error executing statement"
)

func NewMySQLRepository(db *sql.DB) *NivelRepository {
	return &NivelRepository{
		db: db,
	}
}
