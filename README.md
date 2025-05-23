# Votes

A simple votes REST API

## Tech Stack

- **Go**
- **Gin**
- **PostgreSQL**
- **GORM**
- **Go Migrate**
- **Nginx**
- **Docker**
- **Docker Compose**

---

## API

| Endpoint      | Description         |
| ------------- | ------------------- |
| `POST /users` | Register a new user |

---

## Setup & Running the Project

### 1. Clone the repository

```bash
git clone https://github.com/achmadardian/votes.git
```

### 2. Copy environment

```bash
cp .env.example .env
```

- copy `.env.example` to `.env` and update credentials (e.g., DB user/password) **only once** during the first deploy.
- `.env` is included in `.gitignore` to keep sensitive info out of source control.
- subsequent deploys will use the existing `.env` on the server — no need to overwrite unless credentials change.

### 3. Copy database init

```bash
cp initdb/init.sql.example initdb/init.sql
```

- copy `init.sql.example` to `init.sql` and customize it (e.g., update passwords, users) **only once** during the first deploy.
- `init.sql` is typically included in .gitignore to keep sensitive info out of version control.
- subsequent deploys will keep the existing `init.sql` on the server — no need to overwrite unless you change the database setup.

### 4. Install Migrate and Run Migrations

Download and install the [**`migrate` CLI**](https://github.com/golang-migrate/migrate#installation).

Then run:

```bash
migrate -path db/migrations \
  -database "postgres://user:password@localhost:5432/dbname?sslmode=disable" \
  up
```

- make sure PostgreSQL is running and accessible on localhost.
- replace user, password, and dbname with your actual PostgreSQL credentials.
- run this after setting up your database and environment variables, typically only when deploying or updating schema.
- this applies all pending migrations from the `db/migrations` folder.

### 5. Run app

```bash
docker compose up -d
```

## App can be access at

```bash
http://localhost:80/
```
