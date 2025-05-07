# Desafio Fullstack Gazin

Projeto fullstack com o objetivo de gerenciar desenvolvedores e seus respectivos níveis de atuação em uma empresa. A aplicação foi construída com **Go (Backend)** e **Angular (Frontend)**. O projeto foi desenvolvido com uma abordagem purista, focando em simplicidade, organização e reutilização de componentes.

---

## ✨ Tecnologias Utilizadas

| Backend               | Frontend                  |
| :-------------------- | :------------------------ |
| Go                    | Angular                   |
| MySQL                 | TypeScript                |
| Swagger               | SCSS                      |
| Docker/Docker Compose | Componentes Reutilizáveis |
| Hurl                  | Reactive Forms            |
| Makefile              | Serviços REST             |

---

## Execução

A execução do projeto deve ser feito via `Docker/Docker Compose` para evitar problemas de ambiente e dependências (a.k.a "na minha máquina funciona").

### Executando via docker

Certifique-se de que possui instalado: `docker`, `docker-compose`.

Após clonar o repositório e instalar os requisitos, execute os comandos na raiz do projeto:

```bash
cp .example.env .env
docker compose up -d
```

O arquivo `.example.env` possui os valores padrão para a execução do projeto em ambiente de desenvolvimento.

Os serviços estarão disponíveis em:

|                 |                            |
| :-------------- | :------------------------- |
| 🎨 frontend     | http://localhost:4200      |
| 🚀 backend      | http://localhost:8080      |
| 📚 documentação | http://localhost:8080/docs |
| 🛢️ mysql        | http://localhost:3306      |

---

---

## 🧪 Testes

Os testes do backend são realizados via [Hurl](https://hurl.dev/) com arquivos `.hurl` incluídos no diretório `tests/`. Eles cobrem cenários de sucesso e falha para os endpoints de `desenvolvedores` e `niveis`.

Requisições podem ser executadas com o comando: `make request name=caminho_do_arquivo_de_request`  
Testes podem ser executados com o comando: `make test name=caminho_do_arquivo_de_teste`  
O backend pode ser testado completamente com o comando: `make full-test`

---

## 📚 Documentação da API

A documentação da API está disponível via Swagger e roda junto com o serviço de backend em: `http://localhost:8080/docs`

---

## 📁 Migrações

As migrações de banco estão em `./backend/internal/database/migrations/mysql/`, organizadas em arquivos `.up.sql` e `.down.sql`, que são gerados e executados com [golang-migrate](https://github.com/golang-migrate/migrate).

---

## 📦 Estrutura do Projeto

### Backend

A estrutura do backend segue princípios de **Clean Architecture**, separando responsabilidades por camadas independentes:

- `config/`: configurações da aplicação e conexão com o banco.
- `docs/`: arquivos da documentação Swagger gerados pelo Swaggo.
- `internal/`:
  - `database/`: conexão e migrações SQL (via arquivos `.sql`).
  - `entity/`: definição das entidades `Desenvolvedor` e `Nível`.
  - `repository/`: acesso ao banco de dados.
  - `usecase/`: regras de negócio para cada entidade.
  - `presenter/`: transformação de dados para respostas da API.
- `pkg/`: middlewares, utilitários e helpers (CORS, logger, JSON, etc).
- `tests/`: testes automatizados com `hurl`, com cenários de sucesso e erro.
- `main.go`: ponto de entrada da aplicação.

### Frontend

O frontend Angular segue uma arquitetura modular:

- `features/`: módulos para `Desenvolvedores` e `Níveis`, cada um com:
  - `form`: formulário reutilizável em modal.
  - `list`: listagem usando `DataTableComponent`.
  - `services`: comunicação com a API.
  - `models`: modelos de request e response.
- `shared/`: componentes reutilizáveis como:
  - `data-table`: tabela dinâmica com paginação.
  - `form`: formulário dinâmico e standalone.
  - `modal`: controle de exibição de janelas modais.
  - `toast`: sistema de notificações.
- `dashboard/`: container principal da aplicação.

## 🧩 Funcionalidades

### Backend

- CRUD de Desenvolvedores
- CRUD de Níveis (com verificação de vínculo com Desenvolvedores)
- Paginação e ordenação
- Documentação OpenAPI 3.0
- Middleware de CORS, JSON e logger

### Frontend

- Listagem paginada de desenvolvedores e níveis
- Formulários em modal com validação
- Requisições assíncronas com feedback do usuário (toast)
- Interface limpa sem dependências externas

---

## 📌 Considerações

Este projeto foi desenvolvido com foco na clareza do código, responsabilidade única e reutilização de componentes tanto no backend quanto no frontend. Foi evitado o uso de frameworks e bibliotecas além do necessário, priorizando a compreensão e o domínio das bibliotecas padrão do Go e do Angular.
