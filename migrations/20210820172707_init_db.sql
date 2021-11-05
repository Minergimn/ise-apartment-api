-- +goose Up
CREATE TABLE apartment (
  id BIGSERIAL PRIMARY KEY,
  object varchar NOT NULL,
  owner varchar NOT NULL
);

-- +goose Down
DROP TABLE apartment;
