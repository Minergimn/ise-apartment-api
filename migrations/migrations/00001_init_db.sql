-- +goose Up
CREATE TABLE IF NOT EXISTS apartments (
  id BIGSERIAL PRIMARY KEY,
  object varchar NOT NULL,
  owner varchar NOT NULL,
  status int not null
);

CREATE TABLE IF NOT EXISTS apartments_events (
   id BIGSERIAL PRIMARY KEY,
   apartment_id bigint NOT NULL,
   type text NOT NULL,
   status text NOT NULL,
   payload jsonb NULL,
   updated timestamp NULL
);

-- +goose Down
DROP TABLE apartment;
