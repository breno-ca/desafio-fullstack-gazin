package presenter

type Desenvolvedor struct {
	ID              string `json:"id"`
	Nome            string `json:"nome"`
	Sexo            string `json:"sexo"`
	Data_nascimento string `json:"data_nascimento"`
	Hobby           string `json:"hobby"`
	Nivel           struct {
		ID    string `json:"id"`
		Nivel string `json:"nivel"`
	} `json:"nivel"`
}

type DesenvolvedorPaginatedResponse struct {
	Data []Desenvolvedor `json:"data"`
	Meta Meta            `json:"meta"`
}
