package nivel

import "errors"

// CountDesenvolvedoresByNivel counts the number of developers associated with the given id
func (nr *NivelRepository) CountDesenvolvedoresByNivel(id string) (int, error) {
	query := `
		SELECT
			(COUNT(nivel_id))
		FROM
			desenvolvedores
		WHERE
			nivel_id = UUID_TO_BIN(?)
	`

	var count int

	err := nr.db.QueryRow(query, id).Scan(&count)
	if err != nil {
		return 0, errors.New(errorFetchingData)
	}

	return count, nil
}
