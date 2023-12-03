# Dino Diary Server

This project is the backend for the [Dino Diary App](https://github.com/fabim1992/DinoDiary).

## How to Run Locally

### Setup

- Install [Golang](https://go.dev/doc/install);
- Install [Make](https://gnuwin32.sourceforge.net/packages/make.htm);
- Install [Go Migrate Package](https://github.com/golang-migrate/migrate);
- Rename the `.env.example` file to `.env`;

*You also need to have [Docker](https://www.docker.com/) installed to run the development database.*

### Install Dependencies

```sh
# Using Make:
make install
# OR without Make:
go mod tidy
```

### Run

1. Execute postgreSQL database container:

    Start the container:

    ```sh
    docker compose up -d
    ```

    Run migrations:

    ```sh
    make migrate_up
    ```
  
2. Start the server:

    ```sh
    # Using Make:
    make run
    # OR without Make:
    go run ./cmd/server/...
    ```

The server should start at port 3000.

## Routes

[![Run in Insomnia}](https://insomnia.rest/images/run.svg)](https://insomnia.rest/run/?label=Dino%20Diary&uri=https%3A%2F%2Fgithub.com%2Flucasscarioca%2Fdd-server%2Fblob%2Fmain%2Fdocs%2FInsomnia.json)
