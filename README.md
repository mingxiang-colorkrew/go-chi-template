# Golang Playground App

## Migrations

```bash
# run migrations
migrate -source file://db/migrations/ -database "postgres://postgres:postgres@localhost:5432/measure?sslmode=disable" up

# create new migration
migrate create -ext sql -dir db/migrations {migration_name}
```

## Tooling

```bash
brew install golang-migrate
```

## Formatting

```bash
go fmt ./...
```

## Testing

```bash
go test -v ./test/...
```

## VSCode
```bash
go install golang.org/x/tools/gopls@latest
go install -v github.com/go-delve/delve/cmd/dlv@latest
```
