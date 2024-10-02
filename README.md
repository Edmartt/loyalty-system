# Loyalty-System

Loyalty system is a system that implements a solution for the following requirements

- Hexagonal Architecture

- Set campaigns for a store and branches.

- Query campaigns for a store and branches.

- Accumulate points or cashback if the store and branch has a campaign configured.
 
- Redeem points or cashback in a store.


## Project Structure

![](https://github.com/Edmartt/loyalty-system/blob/dev/assets/structure.png)

## ER Diagram

![](https://github.com/Edmartt/loyalty-system/blob/dev/assets/schemeER.png)

## Requirements

- Go 1.19+
- SQLite
- http client: POSTMAN, Insomnia, cURL

### Running Locally

```
git clone https://github.com/Edmartt/loyalty-system.git
```

or ssh instead:

```
git clone git@github.com:Edmartt/loyalty-system.git
```

browse into project directory:

```
cd loyalty-system/
```

download dependencies

```
go mod tidy
```

install tool for migrations

```
go install github.com/rubenv/sql-migrate/...@latest
```

set environment variables following the [.env.example](https://github.com/Edmartt/loyalty-system/blob/dev/.env.example) file for go specific use and run

set environment variables following the [.envrc.example](https://github.com/Edmartt/loyalty-system/blob/dev/.envrc.example) file for migrations config [dbconfig.yml](https://github.com/Edmartt/loyalty-system/blob/dev/dbconfig.yml)

### making migrations

1. create database for the system (do this in your postgres instance)

2. check migrations status. This will load the [migrations/postgres](https://github.com/loyalty-system/blob/dev/migrations/postgres/) files and it will show if migrations are applied or no.

![](https://github.com/Edmartt/loyalty-system/blob/dev/assets/status.png)

```
sql-migrate status
```

3. Apply migrations after checking status

```
sql-migrate up
```

#### note

the above command will apply the migrations in chronological order based on the timestamp. If you want to apply only a specific version you must use the -version flag as follows:

```
sql-migrate up -version 20240929193601
```

4. Run the project

```
go run cmd/main.go
```
