-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE users (
  id serial PRIMARY KEY,
  email varchar(50) NOT NULL,
  password  char(60),
  valid_email boolean NOT NULL DEFAULT FALSE,
  valid_password boolean NOT NULL DEFAULT FALSE,
  UNIQUE(email)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE users;