# Golang Playground App

## Commands

```bash
# run other commands
./commands.sh --help
```

## Setup

Ensure you have the following software installed beforehand
- Docker Compose
- Go (Recommended to install via [asdf](https://asdf-vm.com/))

Run the follow commands to set up the project
```bash
./commands.sh dev_packages:install
./commands.sh init:db
./commands.sh hooks:install
./commands.sh migration:run
go mod download

# after setup is completed. you can start the server via
./commands.sh serve:local
```

## Workflow

### DB changes

1. Create a new migration (refer to `./commands.sh`)
2. Edit the generated `up` and `down` migrations
3. Run the migrations and generate model files (refer to `./commands.sh`)

### API changes

1. Edit the `./openapi.yaml` file with any editor you prefer
    - try to only add 1 endpoint at a time
    - if using certain editors like Stoplight Studio, strip the stoplight tags before commiting the file
2. Autogenerate the OpenAPI go-chi router code (refer to `./commands.sh`)
3. Add handler for new endpoint (trying to run the server will raise compilation errors if you don't)
    - You can generate handler stub via `./commands.sh`
