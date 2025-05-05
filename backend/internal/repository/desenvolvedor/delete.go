package desenvolvedor

import "errors"

// Delete a desenvolvedor from the database
func (dr *DesenvolvedorRepository) Delete(id string) error {
	query := `
		DELETE FROM
		    desenvolvedores
		WHERE
		    id = UUID_TO_BIN(?)
	`
	stmt, err := dr.db.Prepare(query)
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
