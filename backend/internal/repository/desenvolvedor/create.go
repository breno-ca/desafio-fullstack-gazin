package desenvolvedor

import (
	"errors"

	"backend/internal/entity"

	"github.com/go-sql-driver/mysql"
)

// @Summary      Criar nível
// @Description  Cria um novo nível
// @Tags         Níveis
// @Router       /niveis [post]
// @Accept       json
// @Produce      json
// @Success      201      {object}  presenter.Nivel
// @Failure      400      {object}  response.ErrorResponse
// @Failure      500      {object}  response.ErrorResponse
func (dr *DesenvolvedorRepository) Create(desenvolvedor entity.Desenvolvedor) error {
	stmt, err := dr.db.Prepare(`
		INSERT INTO
		    desenvolvedores (id, nivel_id, nome, sexo, data_nascimento, hobby)
		VALUES
		    (UUID_TO_BIN(?), UUID_TO_BIN(?), ?, ?, ?, ?)
	`)
	if err != nil {
		return errors.New(errorPreparingStatement)
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		desenvolvedor.ID,
		desenvolvedor.NivelID,
		desenvolvedor.Nome,
		desenvolvedor.Sexo,
		desenvolvedor.Data_nascimento,
		desenvolvedor.Hobby,
	)
	if err != nil {
		var mysqlError *mysql.MySQLError
		if errors.As(err, &mysqlError) {
			if mysqlError.Number == 1452 {
				return errors.New(errorForeignKeyIDDoesNotExists)
			}
		}

		return errors.New(errorExecutingStatement)
	}

	return nil
}
