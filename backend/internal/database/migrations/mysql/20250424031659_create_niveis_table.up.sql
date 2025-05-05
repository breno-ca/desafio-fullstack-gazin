CREATE TABLE niveis (
    id BINARY(16) NOT NULL PRIMARY KEY,
    nivel VARCHAR(100) NOT NULL
);

CREATE INDEX idx_niveis_id ON niveis(id);
