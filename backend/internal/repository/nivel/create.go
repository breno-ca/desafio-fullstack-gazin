package nivel

import (
	"errors"

	"backend/internal/entity"
)

// Create a new nivel in database
func (nr *NivelRepository) Create(nivel entity.Nivel) error {
	query := `
		INSERT INTO
		    niveis (id, nivel)
		VALUES
		    (UUID_TO_BIN(?), ?)
	`
	stmt, err := nr.db.Prepare(query)
	if err != nil {
		return errors.New(errorPreparingStatement)
	}
	defer stmt.Close()

	_, err = stmt.Exec(nivel.ID, nivel.Nivel)
	if err != nil {
		return errors.New(errorExecutingStatement)
	}

	return nil
}
