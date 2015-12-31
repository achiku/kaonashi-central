
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE account (
  name TEXT PRIMARY KEY
  ,password TEXT NOT NULL
);

CREATE TABLE note (
  id SERIAL PRIMARY KEY
  ,account_name TEXT NOT NULL
  ,title TEXT
  ,body TEXT
  ,created DATE NOT NULL
  ,updated DATE NOT NULL
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE account;
DROP TABLE note;
