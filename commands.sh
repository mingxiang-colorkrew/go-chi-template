#!/bin/bash

cli_help() {
  echo "
  Available Environments(pass as --env=local)
  local, production

  Available commands

  --env={env} serve                               Runs local dev server
  --env={env} migration:run                       Runs migrations and autogenerates DB models
  --env={env} routes:list                         Lists all registered routes

  migration:create {migration_name}   Create a new migration
  openapi:codegen                     Autogenerates OpenAPI Client/Server code 
  format                              Formats output 
  "
}

case "$1" in
  "--env=local")
    echo 'Using local .env.local as ENV'
    set -o allexport
    source .env.local
    set +o allexport
    ;;
  "--env=production")
    echo 'Using production ENV'
    ;;
  "help")
    cli_help
    exit 0
    ;;
  "openapi:codegen")
    echo 'Generating OpenAPI client/server code from oapi/openapi-input.json';
    oapi-codegen -config oapi/codegen.yaml oapi/openapi-input.json
    echo 'Finished generating openapi code';
    exit 0
    ;;
  "format")
    echo 'Formatting all code';
    go fmt ./...
    goimports -w .
    golines -w .
    echo 'Finished formatting all code';
    exit 0
    ;;
  "migration:create")
    echo 'Creating migrations in db/migrations'
    migrate create -ext sql -dir db/migrations "$3"
    echo 'Finished creating migrations';
    exit 0
    ;;
  *)
    echo 'Please provide an env with `./commands.sh --env={env} {command}`'
    cli_help
    exit 1
esac

case "$2" in
  "migration:run")
    echo 'Running migrations from db/migrations';
    migrate -source file://db/migrations/ -database "$DATABASE_URL" up
    echo 'Autogenerating DB models';
    jet -dsn="$DATABASE_URL" -schema=public -path=./db
    # remove reference to migrations table since these should never be interacted with
    rm db/*/*/model/schema_migrations.go
    rm db/*/*/table/schema_migrations.go
    echo 'Finished running migrations';
    ;;
  "serve")
    echo 'Starting server';
    go run .
    ;;
  "routes:list")
    echo 'Listing routes';
    go run . routes:list
    ;;
  *)
    echo 'ERROR: No command provided'
    cli_help
    exit 1
esac
