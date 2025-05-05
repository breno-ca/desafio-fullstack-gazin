### Retorna 404 se não houver nenhum nível cadastrado.
GET {{host}}/api/niveis

HTTP 404
###


### Retorna 400 se o corpo da requisição estiver incorreto.
POST {{host}}/api/niveis
{
  "incorreto": 30
}

HTTP 400
###


### Retorna 400 se o corpo da requisição estiver incorreto.
PUT {{host}}/api/niveis/{{nivel_id}}
{
  "incorreto": 40
}

HTTP 400
###


### Retorna 400 se houver desenvolvedores associados ao nível.
## Insere um nivel e captura o id como dependencia do teste.
POST {{host}}/api/niveis
{
  "nivel": "Trainee"
}

HTTP 201 
[Captures]
nivel_id: jsonpath "$.id"
##

## Insere um desenvolvedor e captura o id como dependencia do teste.
POST {{host}}/api/desenvolvedores
{
  "nome": "desenvolvedor de teste",
  "nivel_id": "{{nivel_id}}",
  "hobby": "Programação",
  "sexo": "M",
  "data_nascimento": "1990-01-01"
}

HTTP 201 
[Captures]
desenvolvedor_id: jsonpath "$.id"
##

# Deleta o nivel
DELETE {{host}}/api/niveis/{{nivel_id}}

HTTP 400
###
