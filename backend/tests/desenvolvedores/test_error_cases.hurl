### Retorna 404 se não houver nenhum desenvolvedor cadastrado.
GET {{host}}/api/desenvolvedores

HTTP 404
###


### Retorna 400 se o corpo da requisição estiver incorreto.
POST {{host}}/api/desenvolvedores
{
  "incorreto": 1,
  "nome": "Nome do Desenvolvedor",
  "faltando": "X",
  "data_nascimento": "1990-01-01",
  "hobby": "Programação"
}

HTTP 400
###


### Retorna 400 se o corpo da requisição estiver incorreto.
PUT {{host}}/api/desenvolvedores/{{desenvolvedor_id}}
{
  "incorreto": 5,
  "name": "Nome do Dev",
  "faltando": "X",
  "data_nascimento": "1990-01-01",
  "hobby": "Programação"
}

HTTP 400
###


### Retorna 400 se houver problemas na remoção. 
DELETE {{host}}/api/desenvolvedores/{{desenvolvedor_id}}

HTTP 400
###
