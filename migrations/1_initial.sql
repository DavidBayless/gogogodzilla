-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE godzillas(id serial, name text, height text);
INSERT INTO godzillas (name, height) VALUES ('Gojira', '200 Giraffes');

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE godzillas;
