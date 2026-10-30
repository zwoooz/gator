-- +goose Up
CREATE TABLE users (
  id UUID PRIMARY KEY,
  name TEXT UNIQUE NOT NULL,
  created_at TEXT NOT NULL,
  updated_at TEXT NOT NULL
);

-- +goose Down
DROP TABLE users;
