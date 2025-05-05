package desenvolvedor

import "database/sql"

type DesenvolvedorRepository struct {
	db *sql.DB
}

const (
	// Error Messages
	errorFetchingData              string = "error fetching data"
	errorPreparingStatement        string = "error preparing statement"
	errorExecutingStatement        string = "error executing statement"
	errorForeignKeyIDDoesNotExists string = "foreign key id does not exist"
)

func NewMySQLRepository(db *sql.DB) *DesenvolvedorRepository {
	return &DesenvolvedorRepository{
		db: db,
	}
}
