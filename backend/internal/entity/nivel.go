package entity

type Nivel struct {
	ID    string `json:"id"`    // ID do nível (UUID)
	Nivel string `json:"nivel"` // Nível do desenvolvedor (Max:100)
}

func (n *Nivel) Validate() error {
	return validate.Struct(struct {
		ID    string `json:"id" validate:"omitempty,uuid"`
		Nivel string `json:"nivel" validate:"required,min=3,max=100"`
	}{
		ID:    n.ID,
		Nivel: n.Nivel,
	})
}
