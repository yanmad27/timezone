# TIMEZONE

## Prerequisites

**Install Encore:**

- **macOS:** `brew install encoredev/tap/encore`
- **Linux:** `curl -L https://encore.dev/install.sh | bash`
- **Windows:** `iwr https://encore.dev/install.ps1 | iex`

## Database

Run the following command to start a local PostgreSQL database:

```bash
docker compose up -d
```

Connection string: `jdbc:postgresql://localhost:5432/postgres`

## Run app locally

```bash
encore run
```

## Debugging

```bash
encore run --debug=break
```

## Local Development Dashboard

While `encore run` is running, open [http://localhost:9400/](http://localhost:9400/) to access Encore's [local developer dashboard](https://encore.dev/docs/go/observability/dev-dash).

Here you can see traces for all requests that you made while using the frontend, see your architecture diagram, and view API documentation in the Service Catalog.
