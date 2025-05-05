package nivel

import (
	"errors"
	"fmt"

	"backend/internal/presenter"
	"backend/pkg/pagination"
	"backend/pkg/util"
)

// Index returns a list of nivel from database
func (nr *NivelRepository) Index(search, by, mode string, page pagination.Pagination) ([]presenter.Nivel, int, error) {
	query := `
		SELECT
		    COUNT(*)
		FROM
		    niveis
		WHERE
		    nivel LIKE ?
	`
	var count int
	err := nr.db.QueryRow(query, search).Scan(&count)
	if err != nil {
		return nil, 0, errors.New(errorExecutingStatement)
	}

	query = `
		SELECT
		    BIN_TO_UUID(niveis.id),
		    niveis.nivel as nivel,
		    COUNT(dev.id) AS qtd_devs
		FROM
		    niveis
		    LEFT JOIN desenvolvedores dev ON niveis.id = dev.nivel_id
		WHERE
		    nivel LIKE ?
		GROUP BY
		    niveis.id,
		    niveis.nivel
		%s
		LIMIT
		    ? OFFSET ?
	`

	switch by {
	case "nivel", "qtd_devs":
		query = fmt.Sprintf(query, util.Order(by, mode))
	default:
		query = fmt.Sprintf(query, "")
	}

	stmt, err := nr.db.Prepare(query)
	if err != nil {
		return nil, 0, errors.New(errorPreparingStatement)
	}
	defer stmt.Close()

	rows, err := stmt.Query(search, page.Limit, page.Offset())
	if err != nil {
		return nil, 0, errors.New(errorExecutingStatement)
	}
	defer rows.Close()

	var niveis []presenter.Nivel
	for rows.Next() {

		var nivel presenter.Nivel
		err = rows.Scan(
			&nivel.ID,
			&nivel.Nivel,
			&nivel.QtdDevs,
		)
		if err != nil {
			return nil, 0, errors.New(errorFetchingData)
		}
		niveis = append(niveis, nivel)
	}

	return niveis, count, nil
}
