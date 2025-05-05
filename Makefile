include .env

## Comandos para iniciar e parar o projeto
launch:
	@echo "Iniciando os serviços do projeto... \n"
	@make env
	@docker compose up -d
	@make migration-up

	@echo "Serviços rodando..."
	@echo " 🎨 frontend   		http://localhost:4200"
	@echo " 🚀 backend   		http://localhost:8080"
	@echo " 📚 documentação 	http://localhost:8080/docs"
	@echo " 🛢️ mysql    		http://localhost:3306"


stop:
	@docker compose down

### Migrações
## Definição das variáveis para acesso ao banco de dados via CLI e execução das migrações
DATABASE_HOST=127.0.0.1
MIGRATIONS_PATH=./backend/internal/database/migrations/mysql
DATA_SOURCE_NAME=mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(DATABASE_HOST):$(MYSQL_PORT))/$(MYSQL_DATABASE)
OPTIONS=charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true
## Definição dos comandos de migração
migration:
	@migrate create -ext=sql -dir=$(MIGRATIONS_PATH) -tz America/Sao_Paulo $(name)

migration-up:
	@migrate -path=$(MIGRATIONS_PATH) -database "$(DATA_SOURCE_NAME)?$(OPTIONS)" -verbose up

migration-down:
	@migrate -path=$(MIGRATIONS_PATH) -database "$(DATA_SOURCE_NAME)?$(OPTIONS)" -verbose down

## Instalação do Go-Migrate com driver MySQL
install-migrate:
	@go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest


 

## Definição dos atalhos para execução de comandos MySQL
MYSQL_BASE_COMMAND=mysql -h ${DATABASE_HOST} -P ${MYSQL_PORT} \
						 -u ${MYSQL_ROOT_USER} -p${MYSQL_ROOT_PASSWORD}
mysql-connection:
	@$(MYSQL_BASE_COMMAND) database

# Exclui o banco de dados e cria novamente
reset-database:
	@$(MYSQL_BASE_COMMAND) -e 'DROP DATABASE `database`; CREATE DATABASE `database`;';
	@make migration-up

# Exclui todos os registros de uma determinada tabela
reset-table:
	@$(MYSQL_BASE_COMMAND) -e 'DELETE FROM `database`.`$(name)`;';

# Executa um script que popula o banco com dados de teste
seed-testdata:
	@$(MYSQL_BASE_COMMAND) < './backend/tests/'$(name)'/testdata/seed.sql'
	@echo " 📝 tabela "$(name)" populada com sucesso!"




## Cria um arquivo .env com base no .example.env
env:
	@cp .example.env .env
	@cp .example.env ./backend/.env




### Testes
## Comando para executar requisições HTTP via CLI
HURL_VARIABLES_FILE=--variables-file=./backend/tests/vars.env
request:
	@bat $(name)
	@hurl --color $(HURL_VARIABLES_FILE) $(name) | jq

## Comando para executar um arquivo de teste
test:
	@hurl --test --color $(HURL_VARIABLES_FILE) $(name)

## Comando para executar todos os testes
full-test:
	@echo "Iniciando os serviços do projeto... \n"
	@docker compose up -d backend --build
	@docker compose up -d database

	@echo "Resetando o banco de dados... \n"
	@make reset-database

	@echo "Testa os casos de erro validando os requisitos de niveis \n"
	make test name=./backend/tests/niveis/test_error_cases.hurl
	@make reset-database

	@echo "Popula o banco com os dados de teste \n"
	@make seed-testdata name=niveis
	@echo "Testa os casos de sucesso validando os requisitos de niveis \n"
	@make test name=./backend/tests/niveis/test_success_cases.hurl
	@make reset-database

	@echo "Testa os casos de erro validando os requisitos de desenvolvedores \n"
	@make test name=./backend/tests/desenvolvedores/test_error_cases.hurl
	@make reset-database

	@echo "Popula o banco com os dados de teste \n"
	@make seed-testdata name=desenvolvedores
	@echo "Testa os casos de sucesso validando os requisitos de desenvolvedores \n"
	@make test name=./backend/tests/desenvolvedores/test_success_cases.hurl
	@make reset-database

	@docker compose down

	@echo " ✅ Todos os testes foram executados com sucesso!"
