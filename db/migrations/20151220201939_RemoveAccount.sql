
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
DROP TABLE account;
ALTER TABLE note DROP account_name;


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
CREATE TABLE account (
  name TEXT PRIMARY KEY
  ,password TEXT NOT NULL
);

ALTER TABLE note ADD account_name;
UPDATE note SET account_name='test';
