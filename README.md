# Booker

## Installation

### Database
Run docker compose to spin-up the database

```console
docker compose up -d
```

TIP: Make sure to use `docker compose` without a dash so you use docker compose v2 instead of v1`

### Build UI
You need to build the UI. The UI is being embed in the backend.

```console
cd web
npm run build
```

### Run the application
```console
go run main.go
```

## Recommendation

I would use Goland from Jetbrains as IDEA.