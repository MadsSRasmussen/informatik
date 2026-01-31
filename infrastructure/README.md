# Infrastructure

This contains information about the infrastructure necessary for facilitating student exercises under the `../exercises/` directory.

## Table of contents

- [Environment variables](#environment-variables)
- [Local development](#local-development)

## Environment variables

The repo contains an `.env.example` that is usefull for establishing a local connection with a MySQL-database, however the example file notably lacks a `MISTAL_API_KEY` environment variable.

## Local development

Running the system locally requires Docker, as the system utilizes docker compose for managing infrastructure.

The following command can be used to run the project:

```bash
docker compose --profile infra --profile api up --build
```

Note that the pure infrastrucutre (the MySQL-database) is assigned a profile "infra", whilst the go-api is assigned a profile "api".
