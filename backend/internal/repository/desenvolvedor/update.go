package desenvolvedor

import (
	"errors"

	"backend/internal/entity"
)

// Update a desenvolvedor in database
func (dr *DesenvolvedorRepository) Update(desenvolvedor entity.Desenvolvedor) error {
	query := `
		UPDATE
		    desenvolvedores
		SET
		    nivel_id = UUID_TO_BIN(?),
		    nome = ?,
		    sexo = ?,
		    data_nascimento = ?,
		    hobby = ?
		WHERE
		    id = UUID_TO_BIN(?)
	`
	stmt, err := dr.db.Prepare(query)
	if err != nil {
		return errors.New(errorPreparingStatement)
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		desenvolvedor.NivelID,
		desenvolvedor.Nome,
		desenvolvedor.Sexo,
		desenvolvedor.Data_nascimento,
		desenvolvedor.Hobby,
		desenvolvedor.ID,
	)
	if err != nil {
		return errors.New(errorExecutingStatement)
	}
	return nil
}
