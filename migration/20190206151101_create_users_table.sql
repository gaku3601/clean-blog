-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE users (
  id serial PRIMARY KEY,
  email varchar(50) NOT NULL,
  password  char(60) NOT NULL,
  UNIQUE(email)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE users;