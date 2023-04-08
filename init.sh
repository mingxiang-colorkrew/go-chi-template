#!/bin/bash

docker-compose up -d

databases=('measure' 'measure_test')

for db in ${databases[@]}; do
  docker-compose exec db psql -U postgres -c "DROP DATABASE IF EXISTS $db;"
  docker-compose exec db psql -U postgres -c "CREATE DATABASE $db;"
done

echo '-------------------------------------------------------------------'

./commands.sh --env=local migration:run
