# Golang Playground App

## Migrations

```bash
# run migrations
migrate -source file://db/migrations/ -database "postgres://postgres:postgres@localhost:5432/measure?sslmode=disable" up

# create new migration
migrate create -ext sql -dir db/migrations {migration_name}
github.com/golang-migrate/migrate/v4
```

## Autogenerate SQL Models

```bash
jet -dsn="postgresql://postgres:postgres@localhost:5432/measure?sslmode=disable" -schema=public -path=./.gen
```

## Autogenerate from OpenAPI Spec

```bash
```

## Tooling

```bash
brew install golang-migrate
go install github.com/go-jet/jet/v2/cmd/jet@latest
go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
```

## Formatting

```bash
go fmt ./...
```

## Testing

```bash
go test -v ./test/...
```

## VSCode DEbugging
```bash
go install golang.org/x/tools/gopls@latest
go install -v github.com/go-delve/delve/cmd/dlv@latest
```
