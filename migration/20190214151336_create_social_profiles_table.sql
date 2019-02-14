-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE social_profiles (
  id serial PRIMARY KEY,
  service varchar(10) NOT NULL,
  uid varchar(50) NOT NULL,
  user_id int REFERENCES users(id) ON DELETE CASCADE NOT NULL,
  UNIQUE (service, uid)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE social_profiles;