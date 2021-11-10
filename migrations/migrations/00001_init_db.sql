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

CREATE INDEX IF NOT EXISTS apartments_object ON apartments (object);
CREATE INDEX IF NOT EXISTS apartments_owner ON apartments (owner);
CREATE INDEX IF NOT EXISTS apartments_events_apartment_id ON apartments_events (apartment_id);
CREATE INDEX IF NOT EXISTS apartments_events_status ON apartments_events (status);

-- +goose Down
DROP INDEX apartments_object;
DROP INDEX apartments_owner;
DROP INDEX apartments_events_apartment_id;
DROP INDEX apartments_events_status;

DROP TABLE apartments;
DROP TABLE apartments_events;
