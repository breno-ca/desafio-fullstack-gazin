CREATE TABLE desenvolvedores (
    id BINARY(16) NOT NULL PRIMARY KEY,
    nivel_id BINARY(16) NOT NULL,
    nome VARCHAR(100) NOT NULL,
    sexo CHAR(1) NOT NULL,
    data_nascimento DATE NOT NULL,
    hobby VARCHAR(100) NOT NULL,
    	CONSTRAINT fk_nivel_id
		FOREIGN KEY(nivel_id)
		REFERENCES niveis(id)
		ON DELETE RESTRICT
);

CREATE INDEX idx_desenvolvedores_id ON desenvolvedores(id);
CREATE INDEX idx_desenvolvedores_nivel_id ON desenvolvedores(nivel_id);
