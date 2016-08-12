-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE animals(id UNIQUEIDENTIFIER DEFAULT NEWID(), name varchar(25) NOT NULL, height varchar(25) NOT NULL);
INSERT INTO animals (id, name, height) VALUES (NEWID(), 'Gojira', '3000 inches');

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE animals;
