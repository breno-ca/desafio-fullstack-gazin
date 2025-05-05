package nivel

import (
	"errors"

	"backend/internal/entity"
)

// Update a nivel in database
func (nr *NivelRepository) Update(nivel entity.Nivel) error {
	query := `
		UPDATE
		    niveis
		SET
		    nivel = ?
		WHERE
		    id = UUID_TO_BIN(?)
	`
	stmt, err := nr.db.Prepare(query)
	if err != nil {
		return errors.New(errorPreparingStatement)
	}
	defer stmt.Close()

	_, err = stmt.Exec(nivel.Nivel, nivel.ID)
	if err != nil {
		return errors.New(errorExecutingStatement)
	}

	return nil
}
