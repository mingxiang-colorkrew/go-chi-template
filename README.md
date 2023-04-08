# Golang Playground App

## Commands

```bash
# initialize DB
./init.sh

# run other commands
./commands.sh --help
```

## Tooling

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
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
