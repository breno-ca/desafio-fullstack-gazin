### Retorna 200 contendo a lista de niveis existentes.
GET {{host}}/api/niveis

HTTP 200
###


### Retorna 200 contendo a lista de niveis existentes.
GET {{host}}/api/niveis/select-options

HTTP 200
###


### Retorna 201 contendo o novo nivel criado.
POST {{host}}/api/niveis
{
  "nivel": "Teste"
}

HTTP 201
[Captures]
nivel_id: jsonpath "$.id"
###


### Retorna 200 contendo o nivel editado.
PUT {{host}}/api/niveis/{{nivel_id}}
{
  "nivel": "teste atualizado"
}

HTTP 200
###


### Retorna 204 se o nivel for deletado.
DELETE {{host}}/api/niveis/{{nivel_id}}

HTTP 204
###
