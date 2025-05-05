package desenvolvedor

import "errors"

func (dr *DesenvolvedorRepository) HasDesenvolvedorByID(id string) (bool, error) {
	query := `
		SELECT
		    (COUNT(id) > 0)
		FROM
		    desenvolvedores
		WHERE
		    id = UUID_TO_BIN(?)
	`

	var hasDesenvolvedor bool

	err := dr.db.QueryRow(query, id).Scan(&hasDesenvolvedor)
	if err != nil {
		return false, errors.New(errorFetchingData)
	}

	return hasDesenvolvedor, nil
}
