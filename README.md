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
