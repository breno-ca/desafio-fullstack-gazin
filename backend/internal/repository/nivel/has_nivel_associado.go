package nivel

import "errors"

// HasNivelAssociado checks if nivel is associated from the database
func (nr *NivelRepository) HasNivelAssociado(id string) (bool, error) {
	query := `
		SELECT 
		    (COUNT(nivel_id) > 0)
		FROM
		    desenvolvedores
		WHERE
		    nivel_id = UUID_TO_BIN(?)
	`

	var hasNivelAssociated bool

	err := nr.db.QueryRow(query, id).Scan(&hasNivelAssociated)
	if err != nil {
		return true, errors.New(errorFetchingData)
	}

	return hasNivelAssociated, nil
}
