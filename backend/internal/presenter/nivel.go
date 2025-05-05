package presenter

type Nivel struct {
	ID      string `json:"id"`
	Nivel   string `json:"nivel"`
	QtdDevs string `json:"qtd_devs"`
}

type NivelSelectOptions struct {
	ID    string `json:"id"`
	Nivel string `json:"nivel"`
}

type NivelPaginatedResponse struct {
	Data []Nivel `json:"data"`
	Meta Meta    `json:"meta"`
}
