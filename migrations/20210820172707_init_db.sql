-- +goose Up
CREATE TABLE apartment (
  id BIGSERIAL PRIMARY KEY,
  foo BIGINT NOT NULL
);

-- +goose Down
DROP TABLE apartment;
