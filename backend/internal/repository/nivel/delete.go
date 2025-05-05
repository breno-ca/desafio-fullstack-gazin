package nivel

import "errors"

// Delete a nivel from the database
func (nr *NivelRepository) Delete(id string) error {
	query := `
		DELETE FROM
			niveis
		WHERE
		    id = UUID_TO_BIN(?)
	`
	stmt, err := nr.db.Prepare(query)
	if err != nil {
		return errors.New(errorPreparingStatement)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return errors.New(errorExecutingStatement)
	}

	return nil
}
