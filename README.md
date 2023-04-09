# Golang Playground App

## Commands

```bash
# initialize DB
./init.sh

# run other commands
./commands.sh --help
```

## Workflow

### DB changes

1. Create a new migration with (refer to `./commands.sh`)
2. Edit the generated `up` and `down` migrations
3. Run the migrations and generate model files (refer to `./commands.sh`)

### API changes

1. Edit the `oapi/openapi-input.json` file with any editor you prefer
2. Autogenerate the OpenAPI go-chi router code (refer to `./commands.sh`)
3. Add handlers to any new routes (trying to run the server will raise compilation errors if you don't)

Refer to `oapi/generated.go` and find `StrictServerInterface` to check list of methods that need to be implemented
```go
// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {

	// (GET /api/v1/tenant)
	GetApiV1Tenant(ctx context.Context, request GetApiV1TenantRequestObject) (GetApiV1TenantResponseObject, error)

	// (POST /api/v1/tenant)
	PostTenant(ctx context.Context, request PostTenantRequestObject) (PostTenantResponseObject, error)
	// Your GET endpoint
	// (GET /api/v1/tenant/{tenantId})
	GetApiV1TenantTenantId(ctx context.Context, request GetApiV1TenantTenantIdRequestObject) (GetApiV1TenantTenantIdResponseObject, error)
	// Your GET endpoint
	// (GET /api/v1/user)
	GetApiV1User(ctx context.Context, request GetApiV1UserRequestObject) (GetApiV1UserResponseObject, error)

	// (POST /api/v1/user)
	PostApiV1User(ctx context.Context, request PostApiV1UserRequestObject) (PostApiV1UserResponseObject, error)
}
```

## Tooling

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
go install github.com/go-jet/jet/v2/cmd/jet@latest
go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
go install golang.org/x/tools/cmd/goimports@latest
```

## Formatting

```bash
go fmt ./...
goimports -w .
golines -w .
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
