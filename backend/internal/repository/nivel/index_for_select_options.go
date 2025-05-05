package nivel

import (
	"errors"

	"backend/internal/presenter"
)

// IndexForSelectOptions returns a list of nivel from database
func (nr *NivelRepository) IndexForSelectOptions() ([]presenter.NivelSelectOptions, int, error) {
	query := `
		SELECT
		    COUNT(id)
		FROM
		    niveis
	`
	var count int
	err := nr.db.QueryRow(query).Scan(&count)
	if err != nil {
		return nil, 0, errors.New(errorExecutingStatement)
	}

	query = `
		SELECT
		    BIN_TO_UUID(id),
		    nivel
		FROM
		    niveis
	`
	stmt, err := nr.db.Prepare(query)
	if err != nil {
		return nil, 0, errors.New(errorPreparingStatement)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, 0, errors.New(errorExecutingStatement)
	}
	defer rows.Close()

	var niveis []presenter.NivelSelectOptions
	for rows.Next() {

		var nivel presenter.NivelSelectOptions
		err = rows.Scan(
			&nivel.ID,
			&nivel.Nivel,
		)
		if err != nil {
			return nil, 0, errors.New(errorFetchingData)
		}
		niveis = append(niveis, nivel)
	}

	return niveis, count, nil
}
