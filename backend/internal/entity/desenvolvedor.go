package entity

type Desenvolvedor struct {
	ID              string `json:"id"`              // ID do desenvolvedor (UUID)
	NivelID         string `json:"nivel_id"`        // ID do nível (UUID)
	Nome            string `json:"nome"`            // Nome do desenvolvedor (Max:100)
	Sexo            string `json:"sexo"`            // Sexo do desenvolvedor (M, F)
	Data_nascimento string `json:"data_nascimento"` // Data de nascimento do desenvolvedor (DD/MM/YYYY)
	Hobby           string `json:"hobby"`           // Hobby do desenvolvedor (Max:100)
}

func (d *Desenvolvedor) Validate() error {
	return validate.Struct(struct {
		ID              string `json:"id" validate:"omitempty,uuid"`
		NivelID         string `json:"nivel_id" validate:"required,uuid"`
		Nome            string `json:"nome" validate:"required,min=3,max=100"`
		Sexo            string `json:"sexo" validate:"required,oneof=M F"`
		Data_nascimento string `json:"data_nascimento" validate:"required"`
		Hobby           string `json:"hobby" validate:"required,min=3,max=100"`
	}{
		ID:              d.ID,
		NivelID:         d.NivelID,
		Nome:            d.Nome,
		Sexo:            d.Sexo,
		Data_nascimento: d.Data_nascimento,
		Hobby:           d.Hobby,
	})
}
