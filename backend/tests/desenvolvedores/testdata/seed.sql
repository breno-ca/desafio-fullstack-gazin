INSERT INTO `database`.`niveis` (id, nivel) VALUES 
(UUID_TO_BIN("0196871f-1584-73bd-ba95-16f8740faa27"), "Trainee");

INSERT INTO `database`.`desenvolvedores` (id, nivel_id, nome, sexo, data_nascimento, hobby) VALUES
(UUID_TO_BIN("0196887f-9d4e-7624-8067-4ec15081da2c"), UUID_TO_BIN("0196871f-1584-73bd-ba95-16f8740faa27"), "desenvolvedor de teste", "M", "1990-01-01", "programação");
