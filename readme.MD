# First GoLang App

First application in GoLang.

## How to run
First, you need to install Go CLI. After, you need to run `go mod tidy` to install dependecies.

Below, contains the scripts to run.

| Name | Command |
|-------|---------|
| run migrations | `go run ./migrate/migrate.go`
| run start | `CompileDaemon -command="./golang-api`"

## Libs
---
- Gin Framework
- Postgres

## Endpoints
---
The file `postman_collection.json` contains all requests.

- `GET /posts` # Get all posts.
- `GET /posts/:id` # Get specific post
- `POST /posts` # Create post
- `PUT /posts/:id` # Update post
- `DELETE /posts/:id` # Delete post
