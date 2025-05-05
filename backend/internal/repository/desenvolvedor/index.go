package desenvolvedor

import (
	"errors"
	"fmt"

	"backend/internal/presenter"
	"backend/pkg/pagination"
	"backend/pkg/util"
)

// Index returns a list of desenvolvedor from database
func (dr *DesenvolvedorRepository) Index(search, by, mode string, page pagination.Pagination) ([]presenter.Desenvolvedor, int, error) {
	query := `
		SELECT
		    COUNT(*)
		FROM
		    desenvolvedores
		WHERE
			nome LIKE ?
	`
	var count int
	err := dr.db.QueryRow(query, search).Scan(&count)
	if err != nil {
		return nil, 0, errors.New(errorExecutingStatement)
	}

	query = `
		SELECT
		    BIN_TO_UUID(dev.id),
		    dev.nome,
		    dev.sexo,
		    dev.data_nascimento,
		    dev.hobby,
		    BIN_TO_UUID(dev.nivel_id),
			niveis.nivel
		FROM
		    desenvolvedores as dev
			INNER JOIN niveis ON niveis.id = dev.nivel_id
		WHERE
			nome LIKE ?
		%s
		LIMIT
			? OFFSET ?
	`

	switch by {
	case "nome", "sexo", "data_nascimento", "hobby", "nivel":
		query = fmt.Sprintf(query, util.Order(by, mode))
	default:
		query = fmt.Sprintf(query, "")
	}

	stmt, err := dr.db.Prepare(query)
	if err != nil {
		return nil, 0, errors.New(errorPreparingStatement)
	}
	defer stmt.Close()

	rows, err := stmt.Query(search, page.Limit, page.Offset())
	if err != nil {
		return nil, 0, errors.New(errorExecutingStatement)
	}
	defer rows.Close()

	var desenvolvedores []presenter.Desenvolvedor
	for rows.Next() {

		var desenvolvedor presenter.Desenvolvedor
		err := rows.Scan(
			&desenvolvedor.ID,
			&desenvolvedor.Nome,
			&desenvolvedor.Sexo,
			&desenvolvedor.Data_nascimento,
			&desenvolvedor.Hobby,
			&desenvolvedor.Nivel.ID,
			&desenvolvedor.Nivel.Nivel,
		)
		if err != nil {
			return nil, 0, errors.New(errorExecutingStatement)
		}
		desenvolvedores = append(desenvolvedores, desenvolvedor)
	}

	return desenvolvedores, count, nil
}
