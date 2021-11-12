#!/bin/sh

goose -dir migrations \
  postgres "user=user password=password host=postgres port=5432 database=apartment_db sslmode=disable" \
  status