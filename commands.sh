#!/bin/bash
#
use_env() {
    echo "Using .env.${1} as ENV"
    set -o allexport
    source ".env.${1}"
    set +o allexport
}

cli_help() {
  echo "
  Available commands

  codegen:handler                       Autogenerates stubs for handlers
  codegen:migration {migration_name}    Autogenerates empty migration files
  codegen:openapi                       Autogenerates OpenAPI Client/Server code 
  dev_packages:install                  Install all dev only packages
  format                                Formats output 
  init:db                               Initializes databases in local docker volumes
  migration:run                         Runs migrations and autogenerates DB models
  routes:list                           Lists all registered routes
  serve:production                      Runs server
  serve:local                           Runs server with .env.local
  hooks:install                         Install git hooks to format code on commit
  "
}

case "${1}" in
  "help")
    cli_help
    exit 0
    ;;

  "codegen:openapi")
    echo 'Generating OpenAPI client/server code from ./openapi.yaml';
    oapi-codegen -config oapi/codegen.yaml ./openapi.yaml
    echo 'Finished generating openapi code';
    exit 0
    ;;

  "codegen:migration")
    migrate create -ext sql -dir db/migrations "${2}"
    exit 0
    ;;

  "codegen:handler")
    echo 'Analyzing oapi.StrictServerInterface for methods that need to be implemented'
    impl 'h *Handler' oapi.StrictServerInterface
    echo 'Please check which methods need to be implement (hint: grep the output)'
    echo '(e.g. ./commands.sh codegen:handler | grep api/v1/user -A 3)'
    exit 0
    ;;

  "dev_packages:install")
    echo 'Installing tools for local development'
    go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
    go install -v github.com/go-delve/delve/cmd/dlv@latest
    go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
    go install github.com/go-jet/jet/v2/cmd/jet@latest
    go install github.com/josharian/impl@latest
    go install golang.org/x/tools/cmd/goimports@latest
    go install golang.org/x/tools/gopls@latest
    go install golang.org/x/tools/cmd/cover@latest
    echo 'All tools installed successfully'

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

  "init:db")
    echo 'Initalize databases';
    docker-compose up -d

    use_env "local"
    docker-compose exec db psql -U postgres -c "DROP DATABASE IF EXISTS ${DATABASE_NAME};"
    docker-compose exec db psql -U postgres -c "CREATE DATABASE ${DATABASE_NAME};"

    use_env "test"
    docker-compose exec db psql -U postgres -c "DROP DATABASE IF EXISTS ${DATABASE_NAME};"
    docker-compose exec db psql -U postgres -c "CREATE DATABASE ${DATABASE_NAME};"
    exit 0
    ;;

  "serve:production")
    echo 'Starting server';
    go run .
    exit 0
    ;;

  "serve:local")
    echo 'Starting server with .env.local';
    use_env "local"
    go run .
    exit 0
    ;;

  "routes:list")
    echo 'Listing routes';
    use_env "local"
    go run . routes:list
    exit 0
    ;;

  "migration:run")
    # run migrations for local db
    echo 'Running migrations from db/migrations';

    echo '------------------------------------------------------------'
    echo 'Migrating local DB';
    use_env "local"
    migrate -source file://db/migrations/ -database "$DATABASE_URL" up

    echo '------------------------------------------------------------'

    echo 'Autogenerating DB models';
    jet -dsn="$DATABASE_URL" -schema=public -path=./db

    echo '------------------------------------------------------------'

    echo 'Migrating test DB';
    use_env "test"
    migrate -source file://db/migrations/ -database "$DATABASE_URL" up
    echo 'Finished running migrations';
    exit 0
    ;;

  "hooks:install")
    # install PHP and npm packages locally so we can run formatting outside docker
    bash ./commands.sh dev_packages:install
    # copy git hook to hooks directory
    cp ./pre-commit ./.git/hooks/pre-commit
    chmod +x .git/hooks/pre-commit
    echo "Copied git hook to ./.git/hooks/pre-commit"
    exit 0
    ;;

  "test")
    use_env "test"
    go test ./test/...
    ;;
    
  *)
    echo 'ERROR: No command provided'
    cli_help
    exit 1
    ;;
esac
