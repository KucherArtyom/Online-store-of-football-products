#!/bin/sh

until psql -h "db" -U "postgres" -c '\q'; do
  echo "Postgres is unavailable - sleeping"
  sleep 1
done

echo "Applying migrations..."
migrate -path ./migrations -database "postgres://postgres:Kuc1804SX@db:5432/footballstore?sslmode=disable" up