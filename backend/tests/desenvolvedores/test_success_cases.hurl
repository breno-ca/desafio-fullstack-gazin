### Retorna 200 contendo a lista de desenvolvedores existentes.
GET {{host}}/api/desenvolvedores

HTTP 200
###


### Retorna 201 contendo o novo desenvolvedor criado.
POST {{host}}/api/desenvolvedores
{
  "nome": "Mario Carrion",
  "nivel_id": "{{nivel_id}}",
  "hobby": "Programação",
  "sexo": "M",
  "data_nascimento": "1990-01-01"
}

HTTP 201
[Captures]
desenvolvedor_id: jsonpath "$.id" 
###


### Retorna 200 contendo o desenvolvedor editado.
PUT {{host}}/api/desenvolvedores/{{desenvolvedor_id}}
{
  "nome": "Mario Carrion",
  "nivel_id": "{{nivel_id}}",
  "hobby": "Tocar Violão",
  "sexo": "M",
  "data_nascimento": "1986-03-08"
}

HTTP 200
###


### Retorna 204 se o desenvolvedor for deletado.
DELETE {{host}}/api/desenvolvedores/{{desenvolvedor_id}}

HTTP 204
###
