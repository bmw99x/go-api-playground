# go-api-playground

## Description

This is a playground for me to learn Go and build a simple API.
I intend on using TDD to incrementally add complexity to the application,
starting with a simple in-memory database, and migrating to a more
production-esque environment with a PostgreSQL database and feature
flags via redis.

## Just commands
Just is a great tool available [here](https://github.com/casey/just) that
allows you to define commands in a `justfile` and run them with `just <command>`.
It doesn't suffer from the same licensing issues as make, and feels a bit more modern
and intuitive. There are several commands to build the application and run it
within local or docker environments.

## Running the application
Either run: `just` after installing just, or use `docker-compose up` to run the application.

## Running tests
Run `just test` to run the tests. This will run the tests on your local machine.

## Swagger documentation
Swagger documentation is available at `http://localhost:8080/swagger/index.html` upon
running the application.

## Configuring feature flags

Edit `flags.yaml` defaultRule if you'd like to change the default behavior of the feature flags.
You can also configure variations by adding rules under `variations` key. For more information,
please see the [documentation](https://gofeatureflag.org/docs/configure_flag/flag_format).

## Generating models

We use [jet](https://github.com/go-jet/jet/v2) to generate models from our database schema.
For example, if we want to use the `public` schema with the postgres user
and postgres password, run the following command to generate representations:
    
```bash
jet "-dsn=postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" -path=./gen --schema=public
```

## Generating migrations
We use atlas to generate migrations and apply them to our database.

Generate initial schema
atlas schema inspect -u "postgres://postgres:postgres@0.0.0.0:5432/postgres?search_path=go_playground&sslmode=disable" > schema.hcl

Generate migration plan
atlas schema apply

Apply migrations
atlas migrate apply -u "postgres://postgres:postgres@0.0.0.0:5432/postgres?search_path=public&sslmode=disable"

Diff between current state and desired state, where schema.hcl is source of truth
atlas schema diff --to "file://schema.hcl" --from "postgres://postgres:postgres@localhost:5432/postgres?search_path=go_playground&sslmode=disable" --dev-url "docker://postgres/15"

Diff between current state and desired state, where postgres is source of truth
atlas schema diff --to "postgres://postgres:postgres@localhost:5432/postgres?search_path=go_playground&sslmode=disable" --from "file://schema.hcl" --dev-url "docker://postgres/15"

Apply state of postgres to schema.hcl
atlas schema apply -u "postgres://postgres:postgres@localhost/postgres?search_path=go_playground&sslmode=disable" --to "file://schema.hcl"

